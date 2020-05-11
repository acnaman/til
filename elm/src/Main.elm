module Main exposing(..)

type LUser = LoggedIn Bool String | Guest

lmessage : LUser -> String
lmessage user = case user of
  LoggedIn isAdmin name ->
    if isAdmin then
      "Hello, " ++ name ++ "(Administrator)."
    else
      "Hello, " ++ name ++ "."
  Guest -> 
    "Please Login."

