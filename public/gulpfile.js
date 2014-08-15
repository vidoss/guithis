var gulp = require('gulp'),
	 rename = require('gulp-rename'),
	 dust = require('gulp-dust'),
	 browserify = require('gulp-browserify');

gulp.task('dust', function() {
	return gulp.src('templates/*.dust')
			.pipe(dust())
			.pipe(gulp.dest('js/tmpl'));
});

gulp.task('browserify', function() {
	return gulp.src('js/main.js')
			.pipe(browserify())
			.pipe(rename('main.min.js'))
			.pipe(gulp.dest('js'));
});

gulp.task('watch', function() {
	gulp.watch('templates/*.dust', ['dust']);
	gulp.watch('js/*.js', ['browserify']);
});
