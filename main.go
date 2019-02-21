package main

import 
(
  "io/ioutil"
  "fmt"
  //"image"
  "github.com/paulmach/go.geojson"
  "github.com/fogleman/gg"
  "time"
)

const (
  width = 400
  height = 400
)

func main() {
  t1 := time.Now()
  rawFeatureCollectionJSON, err := ioutil.ReadFile("./data.geojson")
  if err != nil {
    fmt.Printf("Coulnd't load data.geojson file: %v", err)
    return
  }

  fc := geojson.NewFeatureCollection()
  fc, err = geojson.UnmarshalFeatureCollection(rawFeatureCollectionJSON)
  if err != nil {
    fmt.Printf("Error decoding the data into GeoJSON feature collection: %v", err)
    return
  }

  /*ctx := gg.NewContext(width, height)
  ctx.SetRGB(0, 0, 0)
  ctx.Clear()
  ctx.InvertY()*/
  //img := image.NewRGBA(image.Rect(-width / 2, -height / 2, width / 2, height / 2))
  //img := image.NewRGBA(image.Rect(0, 0, width, height))
  //ctx := gg.NewContextForRGBA(img)
  ctx := gg.NewContext(width, height)
  //ctx.DrawImageAnchored(img, 200, 200, 0.5, 0.5)
  //ctx.InvertY()
  ctx.SetRGB(0, 0, 0) // Paint it, black!
  ctx.Clear()
  
  //ctx.InvertY()

  for _, f := range fc.Features {
    switch {
      case f.Geometry.IsLineString():
        fmt.Println("Detected LineString type of data")
        if ctx, err = DrawLineString(ctx, f); err != nil {
          fmt.Printf("Couldn't handle LineString type: %v", err)
          return
        }
      default:
        fmt.Println("Oops, exotic data type")
        return
    }
  }
  //ctx.DrawImageAnchored(ctx.Image(), 200, 200, 0.5, 0.5)
  err = ctx.SavePNG("image.png")
  if err != nil {
    fmt.Println("Tragic event saving context as PNG: %v", err)
  }

  fmt.Printf("\nProgram finished in %v ", time.Now().Sub(t1))
}

func DrawLineString(c *gg.Context, f *geojson.Feature) (res *gg.Context, err error) {
  c.ClearPath()
  c.SetRGB(0.543, 0, 0)
  c.SetLineWidth(0.5)
  val := f.Geometry.LineString
  /*for i := 1; i < len(val); i++ {
    fmt.Printf("%v\n", val[i - 1])
    c.DrawLine(val[i - 1][0], val[i - 1][1], val[i][0], val[i][1])
  }*/
  for i := 0; i < len(val); i++ {
    c.LineTo(val[i][0] + width / 2, val[i][1] + height / 2)
  }
  //c.SetFillRule(gg.FillRuleEvenOdd) // ?
  c.FillPreserve() // fills the current path with the current color
  c.Stroke()
  c.Fill()

  return c, err
}