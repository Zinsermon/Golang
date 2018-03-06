package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "encoding/xml"
        "html/template"
        "sync"
      )

var waitgroup sync.WaitGroup

func index_handler(writer http.ResponseWriter, request *http.Request){
  // this is format and print on the page
  fmt.Fprintf(writer,"YOLO The Server")
}

func NewsRoutine(c chan News,Locations string){
  defer waitgroup.Done()

  var n News

// before the logic starts
  response, _ :=http.Get(Locations)
  bytes,_ := ioutil.ReadAll(response.Body)
  xml.Unmarshal(bytes,&n)
  response.Body.Close()
  c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

  var s SiteMapIndex
  newsmap := make(map[string]NewsMap)
  // ------------XML-------------------
  response, _ :=http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes,_ := ioutil.ReadAll(response.Body)
  xml.Unmarshal(bytes,&s)
    /// looping through the locations bytes array with range like in python
  // range returns index, value on the SLICE (kind of array)
  news_chan  := make(chan News,50)
  for _,value := range s.Locations{

    waitgroup.Add(1)
    go NewsRoutine(news_chan,value)
    waitgroup.Wait()
    close(news_chan)
}
    for element :=  range news_chan{
      // n is a struct containing 3-arrays of strings , not it self an array
      for index,_ := range element.Keyword{
        newsmap[element.Titles[index]] = NewsMap{element.Keyword[index],element.Locations[index]}

    }
    response.Body.Close()
    }
  // Getting values from n an map of struct type
  // Println by iterating through
  // for idx, data := range newsmap {
  // 		fmt.Println("\n\n\n\n\n",idx)
  // 		fmt.Println("\n",data.Keyword)
  // 		fmt.Println("\n",data.Locations)
  // 	}
	p := NewsAggPage{Title: "Amazing News Aggregator", News:newsmap}
  t, _ := template.ParseFiles("aggregatorfinish.html")
  fmt.Println(t.Execute(w, p))
}

/// -------------------------XML-----------------------
var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

//attributes of the struct should be upper case other wise will get exported
type SiteMapIndex struct{

                        //<sitemap>
                        //   <loc> is inside sitemap tag
  Locations []string `xml:"sitemap>loc"`
}
type News struct{
  Titles []string `xml:"url>news>title"`
  Keyword []string `xml:"url>news>keywords"`
  Locations[]string `xml:"url>loc"`
}
// map doesnot support mutliple values so will make a map of struct type having 2 Values
type NewsMap struct {
  Keyword string
  Locations string
}
type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

func main(){
  // Handler
  http.HandleFunc("/",index_handler)
  http.HandleFunc("/agg/", newsAggHandler)
  /// want functions  to run on the server
  http.ListenAndServe(":8888",nil)
  }
