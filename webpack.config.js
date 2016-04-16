const path = require( 'path' ),
      webpack = require( 'webpack' ),
      paths = {
        build : path.join( __dirname, 'build' ),
        public : path.join( __dirname, 'public' )
      };

module.exports = {
    entry : {
        public : paths.public
    },
    output : {
        path : paths.build,
        filename : 'bundle.js'
    },
    module : {
        loaders : [
            {
                test : /\.jsx?$/,
                exclude : /node_modules/,
                loader : 'babel'
            }
        ]
    },
    devServer : {
        contentBase : paths.build,
        historyApiFallback : true,
        hot : true,
        inline : true,
        progress : true
    },
    plugins: [
        new webpack.HotModuleReplacementPlugin()
    ]
};
