# Local Binary Patterns Histograms (LBPH)

Local Binary Patterns (LBP) is a type of visual descriptor used for classification in computer vision. LBP was first described in 1994 and has since been found to be a powerful feature for texture classification. It has further been determined that when LBP is combined with the Histogram of oriented gradients (HOG) descriptor, it improves the detection performance considerably on some datasets.

You can use go get:

```
go get github.com/kelvins/lbph
```

## Step-by-Step

...

## Important Notes

- The histograms comparison function is using euclidean distance as follows:

$D_{L2} = \sqrt{\sum_{i}\left( h_1(i) - h_2(i) \right) ^2 }$

- It is currently using a fixed radius of 1.

## Usage

...

## Documentation

You can access the full documentation here: [![GoDoc](https://godoc.org/github.com/kelvins/lbph?status.svg)](https://godoc.org/github.com/kelvins/lbph)

## License

This project was created under the **MIT license**. You can read the license here: [![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)

## How to contribute

Feel free to contribute by commenting, suggesting, creating issues or sending pull requests. Any help is welcome.

### Contributing

1. Create an issue (optional)
2. Fork the repo to your Github account
3. Clone the project to your local machine
4. Make your changes
5. Commit your changes (`git commit -am 'Some cool feature'`)
6. Push to the branch (`git push origin master`)
7. Create a new Pull Request
