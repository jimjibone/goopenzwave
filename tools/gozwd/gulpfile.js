// Watch the project, compile and reload the browser on changes.
var gulp        = require('gulp');
var browserify  = require('browserify');
var source      = require('vinyl-source-stream');
// var reactify    = require('reactify');
var imagemin    = require('gulp-imagemin');
var sass        = require('gulp-sass');
var livereload  = require('gulp-livereload');

var sourcemaps = require('gulp-sourcemaps');
var buffer = require('vinyl-buffer');
var watchify = require('watchify');
var babel = require('babelify');

// Source paths:
var src = {
    jsdir:  'assets/js/',
    jsmain: 'main.js',
    js:     ['assets/js/*.js', 'assets/js/*.jsx', 'assets/js/**/*.js', 'assets/js/**/*.jsx'],
    jslibs: ['assets/js/libs/*.js'],
    img:    ['assets/img/*.*', 'assets/img/**/*.*'],
    css:    ['assets/css/*.css', 'assets/css/**/*.css'],
    sassmain: ['assets/css/main.scss'],
    sass:   ['assets/css/*.scss', 'assets/css/**/*.scss'],
    html:   ['assets/*.html']
};

// Destination paths:
var dest = {
    js:     'public/js',
    jslibs: 'public/js/libs',
    img:    'public/img',
    css:    'public/css',
    sass:   'public/css',
    html:   'public/'
};

gulp.task('compile:js', function(){
    // var b = browserify({
    //     debug: true
    // });
    // b.transform(reactify);
    // b.add(src.jsdir+src.jsmain);
    // return b.bundle()
    //     .pipe(source(src.jsmain))
    //     .pipe(gulp.dest(dest.js))
    //     .pipe(livereload());

    var bundler = watchify(browserify(src.jsdir+src.jsmain, { debug: true }).transform(babel));

    return bundler.bundle()
        .on('error', function(err) { console.error(err); this.emit('end'); })
        .pipe(source(src.jsmain))
        .pipe(buffer())
        .pipe(sourcemaps.init({ loadMaps: true }))
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest(dest.js))
        .pipe(livereload());
});

gulp.task('compile:jslibs', function () {
    return gulp.src(src.jslibs)
        .pipe(gulp.dest(dest.jslibs))
        .pipe(livereload());
});

gulp.task('compile:img', function () {
    return gulp.src(src.img)
        .pipe(imagemin())
        .pipe(gulp.dest(dest.img))
        .pipe(livereload());
});

gulp.task('compile:css', function () {
    return gulp.src(src.css)
        .pipe(gulp.dest(dest.css))
        .pipe(livereload());
});

gulp.task('compile:sass', function () {
    return gulp.src('assets/css/main.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest('public/css'))
        .pipe(livereload());
});

gulp.task('compile:html', function () {
    return gulp.src(src.html)
        .pipe(gulp.dest(dest.html))
        .pipe(livereload());
});

// Bundle all compile tasks together.
gulp.task('compile', ['compile:js', 'compile:jslibs', 'compile:img', 'compile:css', 'compile:sass', 'compile:html']);

gulp.task('default', ['compile'], function () {
    // Watch for file changes and rebuild, then use livereload to reload the
    // asset in the browser.
    livereload.listen();
    gulp.watch(src.js,     ['compile:js']);
    gulp.watch(src.jslibs, ['compile:jslibs']);
    gulp.watch(src.img,    ['compile:img']);
    gulp.watch(src.css,    ['compile:css']);
    gulp.watch(src.sass,   ['compile:sass']);
    gulp.watch(src.html,   ['compile:html']);
});
