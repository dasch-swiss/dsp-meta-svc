import * as express from "express";
import * as compression from "compression";

const app = express();

app.use(compression());

app.use("/build", express.static("./public/build"));

app.get("/", (_req, res) => {
  const options

  res.sendFile("index.html", { root: "./public" });
});

const port = 5000;
app.listen(port, () => console.log(`Example app listening on port ${port}`));
