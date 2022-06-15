# Gofiber - Most Common Words
![Coverage](https://img.shields.io/badge/Coverage-95.1%25-brightgreen)

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
[0] a: 14
[1] z: 8
[2] d: 3
[3] y: 2
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
[0] f: 19
[1] z: 8
[2] a: 3
[3] d: 2
[4] y: 1
--- PASS: TestTopN_Sort (0.00s)
--- PASS: TestTopN_Sort/Sorts_an_empty_KV_slice (0.00s)
--- PASS: TestTopN_Sort/Gets_KV_slice_successfully (0.00s)

=== RUN   TestTopN_UpdateKV
=== RUN   TestTopN_UpdateKV/Updates_an_empty_TopKV_slice
[0] a: 3
[1] : 0
[2] : 0
[3] : 0
[4] : 0
=== RUN   TestTopN_UpdateKV/Out_of_range_index_-_not_updated
[0] : 0
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
--- PASS: TestTopN_UpdateKV/Updates_an_empty_TopKV_slice (0.00s)
--- PASS: TestTopN_UpdateKV/Out_of_range_index_-_not_updated (0.00s)
--- PASS: TestTopN_UpdateKV/Updates_the_TopKV_slice_successfully (0.00s)

=== RUN   TestTopN_UpdateVal
=== RUN   TestTopN_UpdateVal/Updates_the_value_in_an_empty_TopKV_slice
=== RUN   TestTopN_UpdateVal/Updates_the_value_successfully
--- PASS: TestTopN_UpdateVal (0.00s)
--- PASS: TestTopN_UpdateVal/Updates_the_value_in_an_empty_TopKV_slice (0.00s)
--- PASS: TestTopN_UpdateVal/Updates_the_value_successfully (0.00s)

=== RUN   Test_insertToGlobalMap
=== RUN   Test_insertToGlobalMap/Inserts_an_empty_map
[0] key: a, value: 1
[1] key: v, value: 16
[2] key: z, value: 3
[3] key: b, value: 7
[4] key: u, value: 5
Length: 5
=== RUN   Test_insertToGlobalMap/Updates_an_existing_word_successfully
[0] key: a, value: 3
[1] key: v, value: 16
[2] key: z, value: 3
[3] key: b, value: 7
[4] key: u, value: 5
[5] key: d, value: 9
Length: 6
=== RUN   Test_insertToGlobalMap/Updates_an_existing_word_successfully#01
[0] key: z, value: 16
[1] key: b, value: 7
[2] key: u, value: 5
[3] key: d, value: 9
[4] key: g, value: 2
[5] key: w, value: 1
[6] key: a, value: 3
[7] key: v, value: 16
Length: 8
--- PASS: Test_insertToGlobalMap (0.00s)
--- PASS: Test_insertToGlobalMap/Inserts_an_empty_map (0.00s)
--- PASS: Test_insertToGlobalMap/Updates_an_existing_word_successfully (0.00s)
--- PASS: Test_insertToGlobalMap/Updates_an_existing_word_successfully#01 (0.00s)

=== RUN   Test_insertToTopN
=== RUN   Test_insertToTopN/Inserts_an_empty_map
[0] a: 14
[1] z: 8
[2] d: 3
[3] y: 2
[4] f: 1
=== RUN   Test_insertToTopN/Does_not_update
[0] a: 14
[1] z: 8
[2] d: 3
[3] y: 2
[4] f: 1
=== RUN   Test_insertToTopN/Updates_an_existing_word_successfully
[0] a: 14
[1] z: 8
[2] d: 5
[3] y: 2
[4] f: 1
=== RUN   Test_insertToTopN/Inserts_a_new_word_successfully
[0] q: 15
[1] a: 14
[2] z: 8
[3] d: 5
[4] y: 2
--- PASS: Test_insertToTopN (0.00s)
--- PASS: Test_insertToTopN/Inserts_an_empty_map (0.00s)
--- PASS: Test_insertToTopN/Does_not_update (0.00s)
--- PASS: Test_insertToTopN/Updates_an_existing_word_successfully (0.00s)
--- PASS: Test_insertToTopN/Inserts_a_new_word_successfully (0.00s)

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

=== RUN   Test_getStrHandler
=== RUN   Test_getStrHandler/Gets_a_string_and_processes_it_successfully
z 24
v 17
q 15
d 15
o 14
=== RUN   Test_getStrHandler/Gets_a_string_and_processes_it_successfully#01
z 32
v 18
q 15
d 15
o 14
=== RUN   Test_getStrHandler/Gets_a_string_and_processes_it_successfully#02
z 32
o 28
v 18
q 15
d 15
=== RUN   Test_getStrHandler/Gets_an_empty_string_successfully
z 32
o 28
v 18
q 15
d 15
--- PASS: Test_getStrHandler (0.00s)
--- PASS: Test_getStrHandler/Gets_a_string_and_processes_it_successfully (0.00s)
--- PASS: Test_getStrHandler/Gets_a_string_and_processes_it_successfully#01 (0.00s)
--- PASS: Test_getStrHandler/Gets_a_string_and_processes_it_successfully#02 (0.00s)
--- PASS: Test_getStrHandler/Gets_an_empty_string_successfully (0.00s)
PASS

coverage: 95.1% of statements in ./...
```
