module Accessor exposing (..)

import Url exposing (Url)
import Url.Parser exposing ((</>), (<?>), s, int, top, map, Parser, oneOf)
import Url.Parser.Query as Q

type Route
    = Top
    | Login 
    | Articles (Maybe String)
    | Article Int
    | ArticleSettings Int

routeParser : Parser (Route -> a) a
routeParser = 
    oneOf
        [ map Top top 
        , map Login (s "login")
        , map Articles (s "articles" <?> Q.string "search")
        , map Article (s "articles" </> int </> s "settings")
        , map ArticleSettings (s "articles" </> int </> s "settings")
        ]

urlToRoute : Url -> Maybe Route
urlToRoute url = 
    Url.Parser.parse routeParser url

-- 以下な感じで書くと Login と判定される
-- Accessor.urlToRoute (Url.Url Url.Https "localhost" (Just 3000) "login" Nothing Nothing)
