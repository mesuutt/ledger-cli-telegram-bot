#!/bin/bash

tmpDir=$(mktemp -d /tmp/termshot-XXXXXXXXXX)
fifoRasterize="${tmpDir}/rasterize.js"
imgFileName="$(date "+termshot-%Y%m%d-%H%M%S.png")"
imgFolder="${HOME}/Pictures"
maxWidth=1280

allOptionsParsed=false

function help() {
	echo
	echo "termshot"
	echo
	echo "This tool takes a \"screenshot\" of the output of a cli command (including prompted messages and typed answers)."
	echo "The output image is cropped to best fit its content but has a maximum width."
	echo
	echo "Options:"
	echo "    --maxWidth|-w [width]       Limit the maximum outputed image's width to [width] pixels."
	echo "                                Exceeding content is wrapped."
	echo "                                Default: 1280"
	echo "    --filename|-f [filename]    The file name to store the resulting image."
	echo "                                Default: \$(date "+termshot-%Y%m%d-%H%M%S.png")"
	echo "    --outputDir|-o [directory]  The folder where the resulting image will be stored."
	echo "                                Default: \${HOME}/Pictures"
	echo "    --help|-h                   Display this help."
	echo
	echo "Examples:"
	echo "    termshot ls -l --color=auto /"
	echo "    termshot --maxWidth 800 rm -i /tmp/test"
	echo
	echo "Author: Pierre Killy <firstName dot lastName at gmail dot com>"
	echo
	echo "Licence: MIT"
	echo
}

if [ "$1" == "" ]
then
	help
	exit 1
fi

while ! $allOptionsParsed
do
	case $1 in
	-w|--maxWidth)
		maxWidth=$2
		shift # past argument
		shift # past value
	;;
	-f|--filename)
		imgFileName="$2"
		shift # past argument
		shift # past value
	;;
	-o|--outputDir)
		imgFolder="$2"
		shift # past argument
		shift # past value
	;;
	-h|--help)
		help
		exit 0
	;;
	*)    # unknown option
		allOptionsParsed=true
	;;
	esac
done

cat > ${fifoRasterize} <<MYEOF
"use strict";
var page = require('webpage').create();
var fs = require('fs');
var pageWidth = ${maxWidth};
var pageHeight = 1280;

page.viewportSize = { width: pageWidth, height: pageHeight };
page.content = fs.read('/dev/stdin');

window.setTimeout(function() {
  page.render('/dev/stdout', { format: 'png' });
  phantom.exit();
}, 100);
MYEOF

imgName="${imgFolder}/${imgFileName}"

ptybandage "$@" \
	| tee /dev/tty \
	| aha --black --word-wrap \
	| phantomjs ${fifoRasterize} \
	| convert -trim - ${imgName}

rm -rf ${tmpDir}
