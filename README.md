# go-fasttext

Serving [fastText](https://github.com/facebookresearch/fastText) model with go

## Dependencies
- [fastText](https://github.com/facebookresearch/fastText)
- cmake

This version still have dependency to libfasttext shared & static libraries and need to be located somewhere in the system in LD_LIBRARY_PATH (usually in /usr/local/lib)

- Install fastText with [cmake](https://github.com/facebookresearch/fastText#building-fasttext-using-cmake). This will install dependencies to your system.

## The Usage

The use of this library is as easy as loading the model and call `Predict` function. Result will be a list of struct of label and probabilities.

```
    m, err := New("dbpedia.ftz")
	result := m.Predict("this is testing", 5, float32(0.01))
```