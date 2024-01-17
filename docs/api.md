# Liste des routes API
(Documentation temporaire)

## /v1/api/homepage

| METHODE    | FONCTION           | DONNEE               |  
|------------|--------------------|----------------------|
| GET        | homepage.Get       |                      |
| GET/{lang} | homepage.GetByLang |                      |
| POST       | homepage.Post      | HomePageModel & lang |
| DELETE     | homepage.Delete    | lang                 |
| PATCH      | homepage.Update    | HomePageModel & lang |

## v1/api/auth:

| METHODE    | FONCTION           | DONNEE    |  
|------------|--------------------|-----------|
| GET        | auth.Auth          | UserModel |

## /v1/api/mail/contact:

| METHODE | FONCTION           | DONNEE            |  
|---------|--------------------|-------------------|
| POST    | auth.Auth          | ContactFormModel  |

