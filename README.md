# merge-ipynb-web

Merge iPython notebooks. Uses the [`merge-ipynb`](https://github.com/crime-analysis/merge-ipynb) package.

## Usage

For example, to append merge `p1.ipynb`, `p2.ipynb`, `p3.ipynb` and save the result in `merged.ipynb`

```bash
$ curl \
> -F "f=@p1.ipynb" \
> -F "f=@p2.ipynb" \
> -F "f=@p3.ipynb" \
> http://merge-ipynb.appspot.com \
> -o merged.ipynb
```

## Running locally

[Follow the Google App Engine for Go setup](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).

Clone this repo. Then from inside the repo:

```
$ goapp serve .
```

Now you can make requests to `http://localhost:8080`!

## License

MIT.