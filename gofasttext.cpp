#include "gofasttext.h"
#include <cstring>

#include "fastText/src/fasttext.h"

using namespace fasttext;

struct membuf : std::streambuf
{
    membuf(char* begin, char* end) {
        this->setg(begin, begin, end);
    }
};

void * GoFastTextInit(char * path)
{
  FastText* ft_ptr = new FastText();
  ft_ptr->loadModel(std::string(path));
  return (void*) ft_ptr;
}

void GoFastTextFree(void * ft)
{
  FastText* ft_ptr = reinterpret_cast<FastText*>(ft);
  delete ft_ptr;
}

go_fast_text_pair_t* GoFastTextPredict(void * ft, char * word, int k, float threshold, int *result_length)
{
  FastText* ft_ptr = reinterpret_cast<FastText*>(ft);
  
  membuf sbuf(word, word + sizeof(word));
  std::istream in(&sbuf);

  std::vector<std::pair<real, std::string>> predictions;
  ft_ptr->predictLine(in, predictions, k, threshold);

  int result_size = predictions.size();

  go_fast_text_pair_t* pairsArray = (go_fast_text_pair_t*) malloc(result_size * sizeof(go_fast_text_pair_t));

  for (uint i = 0; i < uint(predictions.size()); i++){
    const std::string::size_type label_size = predictions[i].second.size();
    pairsArray[i].label = new char[label_size + 1];
    memcpy(pairsArray[i].label, predictions[i].second.c_str(), label_size + 1);
    pairsArray[i].prob = predictions[i].first;
  }

  *result_length = result_size;

  return pairsArray;
}
