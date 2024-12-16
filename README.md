# go_tokenizer_test

## Description

Following list of 4 go-implemented tokenizers were tested based on 
MSR text file.

1. [gojieba](https://github.com/yanyiwu/gojieba)
2. [gse](https://github.com/go-ego/gse/tree/master?tab=readme-ov-file)
3. [jiebago](https://github.com/wangbin/jiebago)
4. [sego](https://github.com/huichen/sego)

All the 4 tokenizer were using default mode with `hmm` set to `true`
## Test result

| metrics | gojieba | gse    | jiebago | sego   |
|---| ------- | ------ | ------- | ------ |
| P | 81.67%  | 83.79% | 81.93%  | 79.64% |
| R | 81.37%  | 79.45% | 81.63%  | 84.20% |
| F1 | 81.52%  | 81.56% | 81.78%  | 81.86% |
| Time | 1.19s   | 4.41   | 6.82    | 1.56   |


## Conclusion
`gojieba` 
1. Top performance with consideration in the trade-off between precision and performance.
2. Implemented in C++, with higher performance and less resource requirement
3. **Do not support cross-platform compile**

`gse`
1. Best precision
2. Implementation is learned from `sego` and `jiebago`



