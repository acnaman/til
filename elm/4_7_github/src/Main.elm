module Main exposing (main)

import Browser
import Browser.Navigation as Nav
import Html exposing (..)
import Html.Attributes exposing (..)
import Http
import Json.Decode as D exposing (Decoder)
import Url
import Url.Builder
import Route exposing (Route)


-- MAIN

main: Program () Model Msg
main = 
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        , onUrlChange = UrlChanged
        , onUrlRequest = UrlRequested
        }

-- MODEL

type alias Model =
    { key : Nav.Key
    , page : Page
    }

type Page 
    = NotFound
    | ErrorPage Http.Error
    | TopPage
    | UserPage (List Repo)
    | RepoPage (List Issue)

init : () -> Url.Url -> Nav.Key -> (Model, Cmd Msg)
init flags url key = 
    Model key TopPage
        |> goTo (Route.parse url)

--UPDATE

type Msg
    = UrlRequested Browser.UrlRequest
    | UrlChanged Url.Url
    | Loaded (Result Http.Error Page)

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model = 
    case msg of
        UrlRequested urlRequest -> 
            case urlRequest of
                Browser.Internal url ->
                    (model, Nav.pushUrl model.key (Url.toString url))

                Browser.External href ->
                    (model, Nav.load href)

        UrlChanged url -> 
            goTo (Route.parse url) model
        
        Loaded result -> 
            ( { model
                | page = 
                    case result of
                        Ok page ->
                            page

                        Err e ->
                            ErrorPage e
              }
            , Cmd.none
            )

goTo : Maybe Route -> Model -> ( Model, Cmd Msg )
goTo maybeRoute model = 
    case maybeRoute of
        Nothing ->
            ( { model | page = NotFound }, Cmd.none )

        Just Route.Top ->
            ( { model | page = TopPage }, Cmd.none )

        Just (Route.User userName) ->
            ( model 
            , Http.get  
                { url = 
                    Url.Builder.crossOrigin "https://api.github.com"
                        ["users", userName, "repos"]
                        []
                
                , expect = 
                    Http.expectJson
                        (Result.map UserPage >> Loaded)
                        reposDecoder
                }
              
            )
        
        Just (Route.Repo userName projectName) ->
            ( model 
            , Http.get  
                { url = 
                    Url.Builder.crossOrigin "https://api.github.com"
                        ["repos", userName, projectName, "issues"]
                        []
                
                , expect = 
                    Http.expectJson
                        (Result.map RepoPage >> Loaded)
                        issuesDecoder
                }
              
            )

-- SUBSCRIPTIONS

subscriptions : Model -> Sub msg
subscriptions _ = 
    Sub.none

-- VIEW

view : Model -> Browser.Document Msg
view model = 
    { title = "My Github Viewer"
    , body = 
        [ a [href "/"]
            [ section [class "hero is-primary"]
                [ div [ class "hero-body"]
                    [ div [class "container"]
                        [ h1 [ class "title" ] [ text "My Github Viewer"] ]
                    ]
                ]
            ]
        

        , case model.page of 
            TopPage ->
                viewTopPage

            NotFound -> 
                viewNotFound
            
            ErrorPage err ->
                viewErrorPage err

            UserPage repoList ->
                viewUserPage repoList
        
            RepoPage issueList ->
                viewRepoPage issueList
        ]
    }

viewErrorPage : Http.Error -> Html msg
viewErrorPage e = 
    case e of
        Http.BadBody message ->
            pre [] [text message]
    
        _ ->
            text (Debug.toString e)

viewNotFound : Html msg
viewNotFound = 
    text "404 Not Found"

viewTopPage : Html msg
viewTopPage = 
    ul [] 
        [ viewLink (Url.Builder.absolute [ "acnaman" ] [])
        , viewLink (Url.Builder.absolute [ "elm" ] [])
        ]

viewLink : String -> Html msg
viewLink path =
    li [] [ a [ href path ] [ text path ] ]

viewUserPage : List Repo -> Html msg
viewUserPage repos = 
    ul [] 
        ( repos 
            |> List.map
                (\repo -> 
                    viewLink ( Url.Builder.absolute [ repo.owner, repo.name ] [])
                )
        )

viewRepoPage : List Issue -> Html msg
viewRepoPage issues = 
    ul [] (List.map viewIssue issues) 

viewIssue : Issue -> Html msg
viewIssue issue = 
    li []
        [ span [] [ text ("[" ++ issue.state ++ "]")]
        , span [] [ text ("#" ++ String.fromInt issue.number)]
        , span [] [ text issue.title ]
        ]

-- UserPage内Repositoryの持つ情報
type alias Repo = 
    { name : String
    , owner : String
    }

-- RepoPage内Issueの持つ情報
type alias Issue = 
    { title : String
    , state : String
    , number : Int
    }

reposDecoder : Decoder (List Repo)
reposDecoder = 
    D.list repoDecoder

repoDecoder : Decoder Repo
repoDecoder = 
    D.map2 Repo
        (D.field "name" D.string)
        (D.at [ "owner", "login" ] D.string)

issuesDecoder : Decoder (List Issue)
issuesDecoder = 
    D.list issueDecoder

issueDecoder : Decoder Issue
issueDecoder =
    D.map3 Issue
        (D.field "title" D.string)
        (D.field "state" D.string)
        (D.field "number" D.int)

