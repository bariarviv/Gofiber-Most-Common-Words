# Gofiber - Most Common Words
![Coverage](https://img.shields.io/badge/Coverage-96%25-brightgreen)

Service with 2 endpoints:
1. **Post:** receives a string containing comma-separated words and processes it.
2. **Get:** returns the N most words with their histogram rank by descending order.

# Unit tests output

```go
=== RUN   TestNewTopN
=== RUN   TestNewTopN/Initializes_the_TopKV_slice_values_to_0
=== RUN   TestNewTopN/Initializes_the_TopKV_slice_values_to_1
--- PASS: TestNewTopN (0.00s)
    --- PASS: TestNewTopN/Initializes_the_TopKV_slice_values_to_0 (0.00s)
    --- PASS: TestNewTopN/Initializes_the_TopKV_slice_values_to_1 (0.00s)

=== RUN   TestTopN_GetKV
=== RUN   TestTopN_GetKV/Gets_empty_KV_slice
=== RUN   TestTopN_GetKV/Gets_KV_slice_successfully
--- PASS: TestTopN_GetKV (0.00s)
    --- PASS: TestTopN_GetKV/Gets_empty_KV_slice (0.00s)
    --- PASS: TestTopN_GetKV/Gets_KV_slice_successfully (0.00s)

=== RUN   TestTopN_GetKey
=== RUN   TestTopN_GetKey/Gets_an_empty_key_from_New_struct
=== RUN   TestTopN_GetKey/Gets_key_successfully
=== RUN   TestTopN_GetKey/Out_of_range_index_-_gets_an_empty_key
--- PASS: TestTopN_GetKey (0.00s)
    --- PASS: TestTopN_GetKey/Gets_an_empty_key_from_New_struct (0.00s)
    --- PASS: TestTopN_GetKey/Gets_key_successfully (0.00s)
    --- PASS: TestTopN_GetKey/Out_of_range_index_-_gets_an_empty_key (0.00s)

=== RUN   TestTopN_GetVal
=== RUN   TestTopN_GetVal/Gets_an_empty_value_from_New_struct
=== RUN   TestTopN_GetVal/Gets_value_successfully
=== RUN   TestTopN_GetVal/Out_of_range_index_-_gets_-1
--- PASS: TestTopN_GetVal (0.00s)
    --- PASS: TestTopN_GetVal/Gets_an_empty_value_from_New_struct (0.00s)
    --- PASS: TestTopN_GetVal/Gets_value_successfully (0.00s)
    --- PASS: TestTopN_GetVal/Out_of_range_index_-_gets_-1 (0.00s)

=== RUN   TestTopN_GetMinMax
=== RUN   TestTopN_GetMinMax/Gets_min_&_max_values_from_empty_TopKV
=== RUN   TestTopN_GetMinMax/Gets_min_&_max_values_successfully
--- PASS: TestTopN_GetMinMax (0.00s)
    --- PASS: TestTopN_GetMinMax/Gets_min_&_max_values_from_empty_TopKV (0.00s)
    --- PASS: TestTopN_GetMinMax/Gets_min_&_max_values_successfully (0.00s)

=== RUN   TestTopN_Len
=== RUN   TestTopN_Len/Gets_len_for_an_empty_KV_slice
=== RUN   TestTopN_Len/Gets_len_successfully
--- PASS: TestTopN_Len (0.00s)
    --- PASS: TestTopN_Len/Gets_len_for_an_empty_KV_slice (0.00s)
    --- PASS: TestTopN_Len/Gets_len_successfully (0.00s)

=== RUN   TestTopN_Print
=== RUN   TestTopN_Print/Prints_an_empty_KV_slice
    [0] : 0
    [1] : 0
    [2] : 0
    [3] : 0
    [4] : 0
=== RUN   TestTopN_Print/Prints_successfully
    [0] a: 5
    [1] d: 3
    [2] y: 3
    [3] z: 1
    [4] f: 1
--- PASS: TestTopN_Print (0.00s)
    --- PASS: TestTopN_Print/Prints_an_empty_KV_slice (0.00s)
    --- PASS: TestTopN_Print/Prints_successfully (0.00s)

=== RUN   TestTopN_Search
=== RUN   TestTopN_Search/Len_of_empty_KV_slice
=== RUN   TestTopN_Search/Success_get_KV_slice
=== RUN   TestTopN_Search/Success_get_KV_slice#01
--- PASS: TestTopN_Search (0.00s)
    --- PASS: TestTopN_Search/Len_of_empty_KV_slice (0.00s)
    --- PASS: TestTopN_Search/Success_get_KV_slice (0.00s)
    --- PASS: TestTopN_Search/Success_get_KV_slice#01 (0.00s)

=== RUN   TestTopN_Sort
=== RUN   TestTopN_Sort/Sorts_an_empty_KV_slice
    [0] : 0
    [1] : 0
    [2] : 0
    [3] : 0
    [4] : 0
=== RUN   TestTopN_Sort/Gets_KV_slice_successfully
    [0] a: 5
    [1] d: 3
    [2] y: 3
    [3] z: 1
    [4] f: 1
--- PASS: TestTopN_Sort (0.00s)
    --- PASS: TestTopN_Sort/Sorts_an_empty_KV_slice (0.00s)
    --- PASS: TestTopN_Sort/Gets_KV_slice_successfully (0.00s)

=== RUN   TestTopN_UpdateKV
=== RUN   TestTopN_UpdateKV/Out_of_range_index_-_not_updated
    [0] : 0
    [1] : 0
    [2] : 0
    [3] : 0
    [4] : 0
=== RUN   TestTopN_UpdateKV/Updates_an_empty_TopKV_slice
    [0] a: 3
    [1] : 0
    [2] : 0
    [3] : 0
    [4] : 0
=== RUN   TestTopN_UpdateKV/Updates_the_TopKV_slice_successfully
    [0] r: 5
    [1] z: 8
    [2] d: 3
    [3] y: 2
    [4] f: 1
--- PASS: TestTopN_UpdateKV (0.00s)
    --- PASS: TestTopN_UpdateKV/Out_of_range_index_-_not_updated (0.00s)
    --- PASS: TestTopN_UpdateKV/Updates_an_empty_TopKV_slice (0.00s)
    --- PASS: TestTopN_UpdateKV/Updates_the_TopKV_slice_successfully (0.00s)

=== RUN   TestTopN_UpdateVal
=== RUN   TestTopN_UpdateVal/Updates_the_value_in_an_empty_TopKV_slice
    [0] : 3
    [1] : 0
    [2] : 0
    [3] : 0
    [4] : 0
=== RUN   TestTopN_UpdateVal/Updates_the_value_successfully
    [0] d: 3
    [1] y: 3
    [2] a: 2
    [3] z: 1
    [4] f: 1
--- PASS: TestTopN_UpdateVal (0.00s)
    --- PASS: TestTopN_UpdateVal/Updates_the_value_in_an_empty_TopKV_slice (0.00s)
    --- PASS: TestTopN_UpdateVal/Updates_the_value_successfully (0.00s)

=== RUN   Test_getMostCommonWordsHandler
=== RUN   Test_getMostCommonWordsHandler/Gets_a_string_and_processes_it_successfully
=== RUN   Test_getMostCommonWordsHandler/Gets_a_string_and_processes_it_successfully#01
=== RUN   Test_getMostCommonWordsHandler/Gets_a_string_and_processes_it_successfully#02
=== RUN   Test_getMostCommonWordsHandler/Gets_an_empty_string_successfully
--- PASS: Test_getMostCommonWordsHandler (0.00s)
    --- PASS: Test_getMostCommonWordsHandler/Gets_a_string_and_processes_it_successfully (0.00s)
    --- PASS: Test_getMostCommonWordsHandler/Gets_a_string_and_processes_it_successfully#01 (0.00s)
    --- PASS: Test_getMostCommonWordsHandler/Gets_a_string_and_processes_it_successfully#02 (0.00s)
    --- PASS: Test_getMostCommonWordsHandler/Gets_an_empty_string_successfully (0.00s)

=== RUN   Test_insertToGlobalMap
=== RUN   Test_insertToGlobalMap/Inserts_an_empty_word_-_not_updated
=== RUN   Test_insertToGlobalMap/Inserts_an_empty_word_-_not_updated/Inserts_an_empty_word_-_not_updated
    [0] key: z, value: 3
    [1] key: b, value: 1
    [2] key: u, value: 5
    [3] key: a, value: 1
    [4] key: v, value: 2
    Length: 5
=== RUN   Test_insertToGlobalMap/Inserts_a_new_word_successfully
=== RUN   Test_insertToGlobalMap/Inserts_a_new_word_successfully/Inserts_a_new_word_successfully
    [0] key: b, value: 1
    [1] key: u, value: 5
    [2] key: t, value: 10
    [3] key: a, value: 1
    [4] key: v, value: 2
    [5] key: z, value: 3
    Length: 6
=== RUN   Test_insertToGlobalMap/Inserts_a_new_word_successfully#01
=== RUN   Test_insertToGlobalMap/Inserts_a_new_word_successfully#01/Inserts_a_new_word_successfully
    [0] key: v, value: 2
    [1] key: z, value: 3
    [2] key: b, value: 1
    [3] key: u, value: 5
    [4] key: t, value: 10
    [5] key: k, value: 4
    [6] key: a, value: 1
    Length: 7
=== RUN   Test_insertToGlobalMap/Updates_an_existing_word_successfully
=== RUN   Test_insertToGlobalMap/Updates_an_existing_word_successfully/Updates_an_existing_word_successfully
    [0] key: t, value: 10
    [1] key: k, value: 4
    [2] key: a, value: 4
    [3] key: v, value: 2
    [4] key: z, value: 3
    [5] key: b, value: 1
    [6] key: u, value: 5
    Length: 7
=== RUN   Test_insertToGlobalMap/Updates_an_existing_word_successfully#01
=== RUN   Test_insertToGlobalMap/Updates_an_existing_word_successfully#01/Updates_an_existing_word_successfully
    [0] key: v, value: 2
    [1] key: z, value: 23
    [2] key: b, value: 1
    [3] key: u, value: 5
    [4] key: t, value: 10
    [5] key: k, value: 4
    [6] key: a, value: 4
    Length: 7
--- PASS: Test_insertToGlobalMap (0.00s)
    --- PASS: Test_insertToGlobalMap/Inserts_an_empty_word_-_not_updated (0.00s)
        --- PASS: Test_insertToGlobalMap/Inserts_an_empty_word_-_not_updated/Inserts_an_empty_word_-_not_updated (0.00s)
    --- PASS: Test_insertToGlobalMap/Inserts_a_new_word_successfully (0.00s)
        --- PASS: Test_insertToGlobalMap/Inserts_a_new_word_successfully/Inserts_a_new_word_successfully (0.00s)
    --- PASS: Test_insertToGlobalMap/Inserts_a_new_word_successfully#01 (0.00s)
        --- PASS: Test_insertToGlobalMap/Inserts_a_new_word_successfully#01/Inserts_a_new_word_successfully (0.00s)
    --- PASS: Test_insertToGlobalMap/Updates_an_existing_word_successfully (0.00s)
        --- PASS: Test_insertToGlobalMap/Updates_an_existing_word_successfully/Updates_an_existing_word_successfully (0.00s)
    --- PASS: Test_insertToGlobalMap/Updates_an_existing_word_successfully#01 (0.00s)
        --- PASS: Test_insertToGlobalMap/Updates_an_existing_word_successfully#01/Updates_an_existing_word_successfully (0.00s)

=== RUN   Test_insertToTopN
=== RUN   Test_insertToTopN/Inserts_an_empty_word_-_not_updated
    [0] a: 5
    [1] d: 3
    [2] y: 3
    [3] z: 1
    [4] f: 1
=== RUN   Test_insertToTopN/Does_not_update
    [0] a: 5
    [1] d: 3
    [2] y: 3
    [3] z: 1
    [4] f: 1
=== RUN   Test_insertToTopN/Updates_an_existing_word_successfully
    [0] a: 5
    [1] y: 4
    [2] d: 3
    [3] z: 1
    [4] f: 1
=== RUN   Test_insertToTopN/Inserts_a_new_word_successfully
    [0] a: 5
    [1] y: 4
    [2] t: 4
    [3] d: 3
    [4] z: 1
--- PASS: Test_insertToTopN (0.00s)
    --- PASS: Test_insertToTopN/Inserts_an_empty_word_-_not_updated (0.00s)
    --- PASS: Test_insertToTopN/Does_not_update (0.00s)
    --- PASS: Test_insertToTopN/Updates_an_existing_word_successfully (0.00s)
    --- PASS: Test_insertToTopN/Inserts_a_new_word_successfully (0.00s)
PASS



=== RUN   Test_getStrHandler
=== RUN   Test_getStrHandler/Gets_a_string_and_processes_it_successfully
    0
    0
    0
    0
    0
    
=== RUN   Test_getStrHandler/Gets_a_string_and_processes_it_successfully#01
    d 5
    a 5
    r 3
    y 1
    t 1
        
=== RUN   Test_getStrHandler/Gets_a_string_and_processes_it_successfully#02
    d 5
    a 5
    z 5
    r 3
    y 1
        
=== RUN   Test_getStrHandler/Gets_an_empty_string_successfully
    d 5
    a 5
    z 5
    o 5
    r 2
        
--- PASS: Test_getStrHandler (0.00s)
    --- PASS: Test_getStrHandler/Gets_a_string_and_processes_it_successfully (0.00s)
    --- PASS: Test_getStrHandler/Gets_a_string_and_processes_it_successfully#01 (0.00s)
    --- PASS: Test_getStrHandler/Gets_a_string_and_processes_it_successfully#02 (0.00s)
    --- PASS: Test_getStrHandler/Gets_an_empty_string_successfully (0.00s)
PASS

coverage: 96% of statements in ./...
```
