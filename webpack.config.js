
const webpack = require( 'webpack' );
const paths   = {
    build  : `${ __dirname }/build`,
    public : `${ __dirname }/public`
};

module.exports = {
    entry : {
        public : `${ paths.public }/index.jsx`
    },
    output : {
        path : paths.build,
        filename : 'bundle.js'
    },
    module : {
        rules : [
            {
                loader  : 'babel',
                test    : /\.jsx?$/,
                exclude : /node_modules/,
                query   : {
                    presets : [ 'es2015', 'react' ]
                }
            },
            {
                loaders : [ 'style', 'css', 'sass' ],
                test   : /\.scss$/
            }
        ]
    }
};
