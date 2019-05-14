const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const CleanWebpackPlugin = require("clean-webpack-plugin");
const webpack = require("webpack");

module.exports = {
  entry: path.join(__dirname, "src", "index.js"),
  output: {
    path: path.join(__dirname, "build"),
    filename: "index.dist.js"
  },
  mode: process.env.NODE_ENV || "development",
  resolve: {
    modules: [path.resolve(__dirname, "src"), "node_modules"]
  },
  devServer: {
    contentBase: path.join(__dirname, "src"),
    disableHostCheck: true,
    host: "localhost"
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ["babel-loader"]
      },
      {
        test: /\.(css|scss)$/,
        use: ["style-loader", "css-loader"]
      },
      {
        test: /\.(svg|png)$/,
        use: "base64-inline-loader"
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.join(__dirname, "src", "index.html")
    }),
    new CleanWebpackPlugin()
  ]
};
