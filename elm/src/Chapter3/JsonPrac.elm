module Chapter3.JsonPrac exposing(foobarHelp, foobar)

import Json.Decode exposing (..)

type FooBar
  = Foo Int String
  | Bar Bool

foobar : Decoder FooBar
foobar =
  field "type" string
    |> andThen foobarHelp

foobarHelp: String -> Decoder FooBar
foobarHelp type_ =
  case type_ of
    "foo" ->
      map2 Foo 
        (field "id" int)
        (field "name" string)

    "bar" -> 
      map Bar
        (field "flag" bool)

    _ ->
      fail "type should be one of [foo, bar]"

type alias Message = { timestamp : Float, data : Value }

messageDecoder : Decoder Message
messageDecoder = 
  map2 Messsage
    (field "timestamp" float)
    (field "data" value)
    
getData : (Decoder a) -> Message -> Maybe a
getData decoder { data } = 
  decodeValue decoder data
    |> Result.toMaybe

import Json.Encode as Encode
tom : Encode.Value
tom = 
  Encode.object
    [ ( "name", Encode.string "Tom" )
    . ( "age, Encode.int 42")
    ]


