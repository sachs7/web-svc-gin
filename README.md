# Steps to build and start the web server

1. docker build --tag sachs7/web-svc-gin:v1.0 .
2. docker run --rm -p 8082:8082 sachs7/web-svc-gin:v1.0
3. docker push sachs7/web-svc-gin:v1.0
4. Use `Minikube` to start the single node cluster
5. kubectl apply -f web-svc-k8s.yaml  (handles both deployment and service)
   1. if want to execute deployment and service files
   2. kubectl apply -f deployment.yaml 
   3. kubectl apply -f service.yaml
6. kubectl describe svc <service name>
7. kubectl cluster-info (to get the IP address)
8. curl http://<IP.from.above>:,nodePort-from-step-5>/albums

# To Run functional tests [as Jobs]
   
   [Test Repo](https://github.com/sachs7/pytest-web-svc-gin)

1. Functional tests are run as `Job` in k8s cluster
2. Exmaple tests - <PATH-TO-FUNCTIONAL-TESTS>/pytest-web-svc-gin
3. kubectl apply -f func-tests.yaml -- to run the job
4. kubectl get pods -n my-func-tests
5. kubectl logs -f goweb-gin-tests-dbz2s -n my-func-tests
============================= test session starts ==============================
platform linux -- Python 3.8.12, pytest-7.0.1, pluggy-1.0.0
rootdir: /app
collected 3 items

test_get_albums.py ...                                                   [100%]

============================== 3 passed in 0.08s ===============================
5. kubectl get jobs -n my-func-tests                     
NAME              COMPLETIONS   DURATION   AGE
goweb-gin-tests   1/1           3s         6m42s
6. kubectl get pods -n my-func-tests
NAME                    READY   STATUS      RESTARTS   AGE
goweb-gin-tests-dbz2s   0/1     Completed   0          119s

# To Run functional tests as a deployment

1. Exmaple tests - <PATH-TO-FUNCTIONAL-TESTS>/pytest-web-svc-gin
2. kubectl apply -f test-using-deployment.yaml
3. Pod will go into crashLoopStage after running the automated tests
4. kubectl logs -f test-84d74c4964-4bj2j                     
============================= test session starts ==============================
platform linux -- Python 3.8.12, pytest-7.0.1, pluggy-1.0.0 -- /usr/local/bin/python
cachedir: .pytest_cache
rootdir: /app
collecting ... http://192.168.99.100
collected 3 items

test_get_albums.py::test_get_albums http://192.168.99.100:31005/albums
200
[{'id': '1', 'title': 'Blue Train', 'artist': 'John Coltrane', 'country': 'United States'}, {'id': '2', 'title': 'Jeru', 'artist': 'Gerry Mulligan', 'country': 'United Kingdom'}, {'id': '3', 'title': 'Sarah Vaughan and Clifford Brown', 'artist': 'Sarah Vaughan', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}]
PASSED
test_get_albums.py::test_get_album_with_id {'id': '2', 'title': 'Jeru', 'artist': 'Gerry Mulligan', 'country': 'United Kingdom'}
PASSED
test_get_albums.py::test_post_create_album {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}
201
[{'id': '1', 'title': 'Blue Train', 'artist': 'John Coltrane', 'country': 'United States'}, {'id': '2', 'title': 'Jeru', 'artist': 'Gerry Mulligan', 'country': 'United Kingdom'}, {'id': '3', 'title': 'Sarah Vaughan and Clifford Brown', 'artist': 'Sarah Vaughan', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}, {'id': '', 'title': 'Ero', 'artist': 'Enigma', 'country': 'India'}]
PASSED

============================== 3 passed in 0.09s ===============================
