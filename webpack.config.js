
const webpack = require( 'webpack' );
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const paths = {
    dist  : `${ __dirname }/dist`,
    public : `${ __dirname }/public`
};

module.exports = {
    entry : {
        public : `${ paths.public }/components/Main/Main.jsx`
    },
    output : {
        path     : paths.dist,
        filename : 'bundle.js'
    },
    module : {
        rules : [
            {
                test    : /\.jsx?$/,
                loader  : 'babel-loader',
                exclude : /node_modules/,
                query   : {
                    presets : [ 'es2015', 'react' ]
                }
            },
            {
                test   : /\.scss$/,
                loader : ExtractTextPlugin.extract({
                    loader : [ 'css-loader', 'sass-loader' ]
                })
            }
        ]
    },
    plugins : [
        new ExtractTextPlugin(`./bundle.css`)
    ]
    
};
