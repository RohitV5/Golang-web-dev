
const express = require('express')
const app = express()
const port = 3000;

app.use(express.json())
app.use(express.urlencoded({extended:true}))

app.get("/",(req,res)=>{
    res.status(200).send("Welcome to express server")
})

app.post("/post",(req,res)=>{
    res.status(200).send(req.body)
})

app.post("/postform",(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(port,()=>{
    console.log(`Running on ${port}`)
})