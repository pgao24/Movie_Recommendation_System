#!/bin/bash
curl -X GET "http://localhost:8080/my_index/_search" -H "Content-Type: application/json" -d '{"query": {"match":{"title":"example"}}}'
