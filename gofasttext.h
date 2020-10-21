// gofasttext.h
#ifdef __cplusplus
extern "C" {
#endif

  typedef struct {
    char* label;
    float prob;                                                                                         
  } go_fast_text_pair_t;

  typedef void* GoFastText;
  GoFastText GoFastTextInit(char * path);
  void GoFastTextFree(void * ft);
  go_fast_text_pair_t* GoFastTextPredict(void * ft, char * word, int k, float threshold, int *result_length);
#ifdef __cplusplus
}
#endif
