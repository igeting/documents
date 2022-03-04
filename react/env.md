# react

## install dotenv-cli
```
npm install --save dotenv-cli
```

## react config file(.env or .env.xxx)
```
REACT_APP_XXX = xxx
```

## react scripts
```
"scripts": {
    "eject": "react-scripts eject",
    "start": "react-scripts start",
    "test": "react-scripts test",
    "build": "react-scripts build",
    "dev": "dotenv -e .env.dev react-scripts start",
    "build-stg": "dotenv -e .env.stg react-scripts build",
    "build-prd": "dotenv -e .env.prd react-scripts build"
}
```