package main

import (
	"log"
	"runtime"
	"time"

	"github.com/daviddengcn/gcse"
	"github.com/daviddengcn/sophie"
)

var (
	cDB *gcse.CrawlerDB
)

func loadPackageUpdateTimes(fpDocs sophie.FsPath) (map[string]time.Time, error) {
	dir := sophie.KVDirInput(fpDocs)
	cnt, err := dir.PartCount()
	if err != nil {
		return nil, err
	}

	pkgUTs := make(map[string]time.Time)

	var pkg sophie.RawString
	var info gcse.DocInfo
	for i := 0; i < cnt; i++ {
		it, err := dir.Iterator(i)
		if err != nil {
			return nil, err
		}
		for {
			if err := it.Next(&pkg, &info); err != nil {
				if err == sophie.EOF {
					break
				}
				return nil, err
			}

			pkgUTs[string(pkg)] = info.LastUpdated
		}
	}
	return pkgUTs, nil
}

func generateCrawlEntries(db *gcse.MemDB, hostFromID func(id string) string,
	out sophie.KVDirOutput) error {
	now := time.Now()
	groups := make(map[string]sophie.CollectCloser)
	count := 0
	if err := db.Iterate(func(id string, val interface{}) error {
		ent, ok := val.(gcse.CrawlingEntry)
		if !ok {
			log.Printf("Wrong entry: %+v", ent)
			return nil
		}

		if ent.Version == gcse.CrawlerVersion &&
			ent.ScheduleTime.After(now) {
			return nil
		}

		host := hostFromID(id)
		c, ok := groups[host]
		if !ok {
			index := len(groups)
			var err error
			c, err = out.Collector(index)
			if err != nil {
				return err
			}
			groups[host] = c
		}

		count++
		return c.Collect(sophie.RawString(id), &ent)
	}); err != nil {
		return err
	}

	for _, c := range groups {
		c.Close()
	}

	log.Printf("%d entries to crawl for folder %v", count, out.Path)
	return nil
}

func syncDatabases() {
	gcse.DumpMemStats()
	log.Printf("Synchronizing databases to disk...")
	if err := cDB.Sync(); err != nil {
		log.Fatalf("cdb.Sync() failed: %v", err)
	}
	gcse.DumpMemStats()
	runtime.GC()
	gcse.DumpMemStats()
}

func main() {
	log.Println("Running tocrawl tool, to generate crawling list")
	// Load CrawlerDB
	cDB = gcse.LoadCrawlerDB()

	if gcse.CrawlGithubUpdate || gcse.CrawlByGodocApi {
		// load pkgUTs
		pkgUTs, err := loadPackageUpdateTimes(
			sophie.LocalFsPath(gcse.DocsDBPath.S()))
		if err != nil {
			log.Fatalf("loadPackageUpdateTimes failed: %v", err)
		}
		
		if gcse.CrawlGithubUpdate {
			touchByGithubUpdates(pkgUTs)
		}
		
		if gcse.CrawlByGodocApi {
			httpClient := gcse.GenHttpClient("")
			pkgs, err := gcse.FetchAllPackagesInGodoc(httpClient)
			if err != nil {
				log.Fatalf("FetchAllPackagesInGodoc failed: %v", err)
			}
			log.Printf("FetchAllPackagesInGodoc returns %d entries", len(pkgs))
			for _, pkg := range pkgs {
				cDB.AppendPackage(pkg, func(pkg string) bool {
					_, ok := pkgUTs[pkg]
					return ok
				})
			}
		}
		syncDatabases()
	}
	

	log.Printf("Package DB: %d entries", cDB.PackageDB.Count())
	log.Printf("Person DB: %d entries", cDB.PersonDB.Count())

	pathToCrawl := gcse.DataRoot.Join(gcse.FnToCrawl)

	kvPackage := sophie.KVDirOutput(sophie.LocalFsPath(
		pathToCrawl.Join(gcse.FnPackage).S()))
	kvPackage.Clean()
	if err := generateCrawlEntries(cDB.PackageDB, gcse.HostOfPackage, kvPackage); err != nil {
		log.Fatalf("generateCrawlEntries %v failed: %v", kvPackage.Path, err)
	}

	kvPerson := sophie.KVDirOutput{
		Fs:   sophie.LocalFS,
		Path: pathToCrawl.Join(gcse.FnPerson).S(),
	}
	kvPerson.Clean()
	if err := generateCrawlEntries(cDB.PersonDB, func(id string) string {
		site, _ := gcse.ParsePersonId(id)
		return site
	}, kvPerson); err != nil {
		log.Fatalf("generateCrawlEntries %v failed: %v", kvPerson.Path, err)
	}
}
