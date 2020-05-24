module RouteTest exposing (..)

import Expect exposing (Expectation)
import Test exposing (..)
import Url exposing (Url)

import Route exposing (Route, parse)

suite : Test
suite = 
    describe "Route"
        [ testParse "should parse top" "/" (Just Route.Top)
        , testParse "should parse User" "/user01" (Just (Route.User "user01"))
        , testParse "should parse Repository" "/user01/hogeproject" (Just (Route.Repo "user01" "hogeproject"))
        , testParse "should parse invalid" "/foo/bar/baz" (Nothing)
        ]

testParse : String -> String -> Maybe Route -> Test
testParse name path expectedRoute = 
    test name <|
        \_ -> 
            Url.fromString("http://example.com" ++ path)
                |> Maybe.andThen Route.parse
                |> Expect.equal expectedRoute

