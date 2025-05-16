# GONVERTER
Gonverter converts an image (JPG, PNG or JPEG) to a PDF doing it for free and without any kind of annoying ads.

## How to use
1. Clone gonverter and gonverter-web repository.
2. To Run gonverter: `cd gonverter` and then `go run ./...`
3. To run gonverter-web: `cd gonverter-web` then `bun dev`
4. Upload an image using the UI and click download to get the PDF.

## Details
* API and all of the logic is written in Go.
* Gonverter or Gonverter-web does not saves your images or pdfs in any database.
* If the image doesn't fit in an A4 page (210 x 297 milimmeters) then it's resized.
* The only external dependency is the [go-pdf/fpdf](https://codeberg.org/go-pdf/fpdf) which is used to manage all the pdf related code.
* **Gonverter** is related to [gonverter-web](https://github.com/Alastair7/gonverter-web).
* This project does not contains tests because I do not intend to deploy it for production use. It's just a side-project and I had a lot of fun coding it.

## [GONVERTER-WEB](https://github.com/Alastair7/gonverter-web)
Gonverter-Web it's just the web app that uses gonverter's API.
It provides a UI to upload an image and convert it to PDF.

## Details
* Web App is built with TS, React and Tailwind and Bun.
