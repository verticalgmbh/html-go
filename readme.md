## html-go

This package provides an alternative way to parse html data which is not html5 compliant.
It is mainly used to process template html and write it again so don't expect a fully featured parser.

The built in html package of golang sadly is unable to import html files correctly which contain tags not being part
of the html5 specification. While it does not output any errors it just places tags at whatever location it deems appropriate.

The scope of this package is to be able to import these files while preserving the original structure and then output valid html
again after processing some elements.

### Usage

#### Parsing
```
doc,err:=html.Parse(<any io.Reader>)
```

Now **doc** will contain the parsed html document if no error occured while parsing the data. **err** contains any parsing error.
When parsing was successful you can traverse **doc.HTML** to process the HTML tree.

#### Writing
```
html.Write(<*Document>, <io.Writer>)
```

Writes the document to a writer