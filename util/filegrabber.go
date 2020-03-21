package util

import (

	"time"
	 "net/http"
	 "log"
	 "io/ioutil"
	"github.com/cavaliercoder/grab"
)

func ReadURLFile(url string) string {

	resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	return string(data)

}
func FetchFile(dir string, url string,file string) string {

	fname:=dir+"/"+file
  if FileExists(fname) {
    Info("image file %s already exists",file)
    return fname
  }
	client := grab.NewClient()
	req, _ := grab.NewRequest(dir, url+"/"+file)

	// start download
	Info("Downloading %v...\n", req.URL())
	resp := client.Do(req)

	// start UI loop
	t := time.NewTicker(5000 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			Info("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
    Fail("Download failed: %v\n", err)
	}

Info("Download saved to %v \n", resp.Filename)

return fname

}
