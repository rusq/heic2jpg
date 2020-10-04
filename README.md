# HEIC to JPG converter

Quick and dirty heic2jpg converter that currenly runs only on osX.

It converts a single file or walks through the directory (possibly recusively)
converting all heic files over there.  Optionally it can remove the source
files.

To install:

    go get github.com/rusq/heic2jpg
    go install github.com/rusq/heic2jpg/cmd/heic2jpg

## Usage examples

    heic2jpg <filename.heic>

OR:

    heic2jpg -d <directory>

Recursive directory with JPEG quality 90 (if not specified, default is 70):

    heic2jpg -d <directory> -r -q=90

Delete source files:

    heic2jpg -d <directory> -rm

