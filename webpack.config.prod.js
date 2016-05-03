
const webpack = require( 'webpack' ),
      paths   = {
        build  : `${ __dirname }/build`,
        public : `${ __dirname }/public`
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
                loader  : 'babel',
                test    : /\.jsx?$/,
                exclude : /node_modules/
            },
            {
                loader : 'style!css!less',
                test   : /\.less$/
            },
            {
                loader : 'url-loader',
                test : /\.png$/
            }
        ]
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env' : {
                'NODE_ENV' : JSON.stringify( 'production' )
            }
        })
    ]
};
