const path = require("path")
const gulp = require("gulp")
const sass = require("gulp-sass")(require("node-sass"))
const webpack = require("webpack-stream")
const webpackConfig = require("./webpack.config")

const sassSrcPath = path.resolve(__dirname, "scss")
const jsSrcPath = path.resolve(__dirname, "scripts")
const viewPath = path.resolve(__dirname, "../views")
const outputDir = path.resolve(__dirname, ".")
const scssFilesTarget = path.join(sassSrcPath, "*.scss")
const jsFilesTarget = path.resolve(__dirname, "scripts/app.js")
const allJsWatchFiles = path.resolve(jsSrcPath, "**/*")

gulp.task("build:css", function() {
    return gulp.src(scssFilesTarget)
        .pipe(sass().on("error", sass.logError))
        .pipe(gulp.dest(outputDir))
})

gulp.task("build:js", function (){
    return  gulp.src(jsFilesTarget)
        .pipe(webpack(webpackConfig))
        .pipe(gulp.dest(outputDir))
})

gulp.task("watch", function() {
    gulp.watch(path.join(viewPath, '../*.html'))
    gulp.watch(scssFilesTarget, gulp.series("build:css"))
    gulp.watch(allJsWatchFiles, gulp.series("build:js"))
})
