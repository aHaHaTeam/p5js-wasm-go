# p5js-wasm-go

It is a Golang library with bindings to use p5js in the browser from webassembly.

## Features

Here is a list of features that have already been implemented:

#
### Environment
<details>
<summary>✅ Completed</summary>

✅`describe()`

✅`describeElement()`

✅`textOutput()`

✅`gridOutput()`

✅`print()`

✅`frameCount`

✅`deltaTime`

✅`focused`

✅`cursor()`

✅`frameRate()`

✅`getTargetFrameRate()`

✅`noCursor()`

✅`webglVersion`

✅`displayWidth`

✅`displayHeight`

✅`windowWidth`

✅`windowHeight`

✅`windowResized()`

✅`width`

✅`height`

✅`fullscreen()`

✅`pixelDensity()`

✅`displayDensity()`

✅`getURL()`

✅`getURLPath()`

✅`getURLParams()`

</details>

#
### Color
#### Creating & Reading
<details>
<summary>✴️ In progress (10/11)</summary>

✅`alpha()`

✅`blue()`

✅`brightness()`

✅`color()`

✅`green()`

✅`hue()`

✅`lerpColor()`

✅`lightness()`

✅`red()`

✅`saturation()`

⬜`p5.Color`

</details>

#### Setting
<details>

<summary>✅ Completed</summary>

✅`beginClip()`

✅`endClip()`

✅`clip()`

✅`background()`

✅`clear()`

✅`colorMode()`

✅`fill()`

✅`noFill()`

✅`noStroke()`

✅`stroke()`

✅`erase()`

✅`noErase()`

</details>

#
### Shape
#### 2D Primitives
<details>
<summary>✅ Completed</summary>

✅`arc()`

✅`ellipse()`

✅`circle()`

✅`line()`

✅`point()`

✅`quad()`

✅`rect()`

✅`square()`

✅`triangle()`
</details>

#### Attributes
<details>
<summary>✅ Completed</summary>

✅`ellipseMode()`

✅`noSmooth()`

✅`rectMode()`

✅`smooth()`

✅`strokeCap()`

✅`strokeJoin()`

✅`strokeWeight()`
</details>

#### Curves
<details>
<summary>⬜ In progress (0/9)</summary>

⬜`bezier()`

⬜`bezierDetail()`

⬜`bezierPoint()`

⬜`bezierTangent()`

⬜`curve()`

⬜`curveDetail()`

⬜`curveTightness()`

⬜`curvePoint()`

⬜`curveTangent()`
</details>

#### Vertex
<details>
<summary>⬜ In progress (0/9)</summary>

⬜`beginContour()`

⬜`beginShape()`

⬜`bezierVertex()`

⬜`curveVertex()`

⬜`endContour()`

⬜`endShape()`

⬜`quadraticVertex()`

⬜`vertex()`

⬜`normal()`
</details>

#### 3D Primitives
<details>
<summary>⬜ In progress (0/12)</summary>
  
⬜`beginGeometry()`

⬜`endGeometry()`

⬜`buildGeometry()`

⬜`freeGeometry()`

⬜`plane()`

⬜`box()`

⬜`sphere()`

⬜`cylinder()`

⬜`cone()`

⬜`ellipsoid()`

⬜`torus()`

⬜`p5.Geometry`
</details>

#### 3D Models
<details>
<summary>⬜ In progress (0/2)</summary>

⬜`loadModel()`

⬜`model()`
</details>

#
### Constants
<details>
<summary>✅ Completed</summary>
  
✅`HALF_PI`

✅`PI`

✅`QUARTER_PI`

✅`TAU`

✅`TWO_PI`

✅`DEGREES`

✅`RADIANS`
</details>

#
### Structure
<details>
<summary>✴️ In progress (10/12)</summary>

✅`preload()`

✅`setup()`

✅`draw()`

✅`remove()`

⬜`disableFriendlyErrors`

✅`noLoop()`

✅`loop()`

✅`isLooping()`

✅`push()`

✅`pop()`

✅`redraw()`

⬜`p5()`
</details>

#
### DOM
<details>
<summary>⬜ In progress (0/25)</summary>

⬜`p5.Element`

⬜`select()`

⬜`selectAll()`

⬜`removeElements()`

⬜`changed()`

⬜`input()`

⬜`createDiv()`

⬜`createP()`

⬜`createSpan()`

⬜`createImg()`

⬜`createA()`

⬜`createSlider()`

⬜`createButton()`

⬜`createCheckbox()`

⬜`createSelect()`

⬜`createRadio()`

⬜`createColorPicker()`

⬜`createInput()`

⬜`createFileInput()`

⬜`createVideo()`

⬜`createAudio()`

⬜`createCapture()`

⬜`createElement()`

⬜`p5.MediaElement`

⬜`p5.File`
</details>

#
### Rendering
<details>
<summary>✴️ In progress (9/11)</summary>

⬜`p5.Graphics`

✅`createCanvas()`

✅`resizeCanvas()`

✅`noCanvas()`

✅`createGraphics()`

✅`createFramebuffer()`

✅`clearDepth()`

✅`blendMode()`

✅`drawingContext`

✅`setAttributes()`

⬜`p5.Framebuffer`
</details>

#
### Foundation
<details>
<summary>⬜ In progress (0/12)</summary>
  
⬜`let`

⬜`if`

⬜`function`

⬜`Boolean`

⬜`String`

⬜`Number`

⬜`Object`

⬜`Array`

⬜`class`

⬜`for`

⬜`while`

⬜`console`
</details>

#
### Transform
<details>
<summary>⬜ In progress (0/10)</summary>

⬜`applyMatrix()`

⬜`resetMatrix()`

⬜`rotate()`

⬜`rotateX()`

⬜`rotateY()`

⬜`rotateZ()`

⬜`scale()`

⬜`shearX()`

⬜`shearY()`

⬜`translate()`
</details>

#
### Data
#### LocalStorage
<details>
<summary>⬜ In progress (0/4)</summary>

⬜`storeItem()`

⬜`getItem()`

⬜`clearStorage()`

⬜`removeItem()`
</details>

#### Dictionary
<details>
<summary>⬜ In progress (0/4)</summary>

⬜`createStringDict()`

⬜`createNumberDict()`

⬜`p5.TypedDict`

⬜`p5.NumberDict`
</details>

#### Array Functions
<details>
<summary>⬜ In progress (0/9)</summary>

⬜`append()`

⬜`arrayCopy()`

⬜`concat()`

⬜`reverse()`

⬜`shorten()`

⬜`shuffle()`

⬜`sort()`

⬜`splice()`

⬜`subset()`
</details>

#### Conversion
<details>
<summary>⬜ In progress (0/9)</summary>

⬜`float()`

⬜`int()`

⬜`str()`

⬜`boolean()`

⬜`byte()`

⬜`char()`

⬜`unchar()`

⬜`hex()`

⬜`unhex()`
</details>

#### String Functions
<details>
<summary>⬜ In progress (0/10)</summary>

⬜`join()`

⬜`match()`

⬜`matchAll()`

⬜`nf()`

⬜`nfc()`

⬜`nfp()`

⬜`nfs()`

⬜`split()`

⬜`splitTokens()`

⬜`trim()`
</details>

#
### Events
#### Acceleration
<details>
<summary>⬜ In progress (0/38)</summary>

⬜`deviceOrientation`

⬜`accelerationX`

⬜`accelerationY`

⬜`accelerationZ`

⬜`pAccelerationX`

⬜`pAccelerationY`

⬜`pAccelerationZ`

⬜`rotationX`

⬜`rotationY`

⬜`rotationZ`

⬜`pRotationX`

⬜`pRotationY`

⬜`pRotationZ`

⬜`turnAxis`

⬜`setMoveThreshold()`

⬜`setShakeThreshold()`

⬜`deviceMoved()`

⬜`deviceTurned()`

⬜`deviceShaken()`
</details>

#### Keyboard
<details>
<summary>✅ Completed</summary>

✅`keyIsPressed`

✅`key`

✅`keyCode`

✅`keyPressed()`

✅`keyReleased()`

✅`keyTyped()`

✅`keyIsDown()`
</details>

#### Mouse
<details>
<summary>✅ Completed</summary>

✅`movedX`

✅`movedY`

✅`mouseX`

✅`mouseY`

✅`pmouseX`

✅`pmouseY`

✅`winMouseX`

✅`winMouseY`

✅`pwinMouseX`

✅`pwinMouseY`

✅`mouseButton`

✅`mouseIsPressed`

✅`mouseMoved()`

✅`mouseDragged()`

✅`mousePressed()`

✅`mouseReleased()`

✅`mouseClicked()`

✅`doubleClicked()`

✅`mouseWheel()`

✅`requestPointerLock()`

✅`exitPointerLock()`
</details>

#### Touch
<details>
<summary>⬜ In progress (0/4)</summary>

⬜`touches`

⬜`touchStarted()`

⬜`touchMoved()`

⬜`touchEnded()`
</details>

#
### Image
#### Image
<details>
<summary>⬜ In progress (0/4)</summary>

⬜`createImage()`

⬜`saveCanvas()`

⬜`saveFrames()`

⬜`p5.Image`
</details>

#### Loading & Displaying
<details>
<summary>⬜ In progress (0/6)</summary>

⬜`loadImage()`

⬜`saveGif()`

⬜`image()`

⬜`tint()`

⬜`noTint()`

⬜`imageMode()`
</details>

#### Pixels
<details>
<summary>⬜ In progress (0/8)</summary>

⬜`pixels`

⬜`blend()`

⬜`copy()`

⬜`filter()`

⬜`get()`

⬜`loadPixels()`

⬜`set()`

⬜`updatePixels()`
</details>

#
### IO
#### Input
<details>
<summary>⬜ In progress (0/9)</summary>

⬜`loadJSON()`

⬜`loadStrings()`

⬜`loadTable()`

⬜`loadXML()`

⬜`loadBytes()`

⬜`httpGet()`

⬜`httpPost()`

⬜`httpDo()`

⬜`p5.XML`
</details>

#### Output
<details>
<summary>⬜ In progress (0/6)</summary>

⬜`createWriter()`

⬜`p5.PrintWriter`

⬜`save()`

⬜`saveJSON()`

⬜`saveStrings()`

⬜`saveTable()`
</details>

#### Table
<details>
<summary>⬜ In progress (0/2)</summary>

⬜`p5.Table`

⬜`p5.TableRow`
</details>

#### Time & Date
<details>
<summary>⬜ In progress (0/7)</summary>

⬜`day()`

⬜`hour()`

⬜`minute()`

⬜`millis()`

⬜`month()`

⬜`second()`

⬜`year()`
</details>

#
### Math
#### Calculation
<details>
<summary>⬜ In progress (0/18)</summary>

⬜`abs()`

⬜`ceil()`

⬜`constrain()`

⬜`dist()`

⬜`exp()`

⬜`floor()`

⬜`lerp()`

⬜`log()`

⬜`mag()`

⬜`map()`

⬜`max()`

⬜`min()`

⬜`norm()`

⬜`pow()`

⬜`round()`

⬜`sq()`

⬜`sqrt()`

⬜`fract()`
</details>

#### Vector
<details>
<summary>⬜ In progress (0/2)</summary>

⬜`createVector()`

⬜`p5.Vector`
</details>

#### Noise
<details>
<summary>⬜ In progress (0/3)</summary>

⬜`noise()`

⬜`noiseDetail()`

⬜`noiseSeed()`
</details>

#### Random
<details>
<summary>⬜ In progress (0/18)</summary>

⬜`randomSeed()`

⬜`random()`

⬜`randomGaussian()`
</details>

#### Trigonometry
<details>
<summary>⬜ In progress (0/18)</summary>

⬜`acos()`

⬜`asin()`

⬜`atan()`

⬜`atan2()`

⬜`cos()`

⬜`sin()`

⬜`tan()`

⬜`degrees()`

⬜`radians()`

⬜`angleMode()`
</details>

#
### Typography
#### Attributes
<details>
<summary>✅ Completed</summary>

✅`textAlign()`

✅`textLeading()`

✅`textSize()`

✅`textStyle()`

✅`textWidth()`

✅`textAscent()`

✅`textDescent()`

✅`textWrap()`
</details>

#### Loading & Displaying
<details>
<summary>✴️ In progress (3/4)</summary>

✅`loadFont()`

✅`text()`

✅`textFont()`

⬜`p5.Font`
</details>

#
### 3D
#### Interaction
<details>
<summary>⬜ In progress (0/3)</summary>

⬜`orbitControl()`

⬜`debugMode()`

⬜`noDebugMode()`
</details>

#### Lights
<details>
<summary>⬜ In progress (0/10)</summary>

⬜`ambientLight()`

⬜`specularColor()`

⬜`directionalLight()`

⬜`pointLight()`

⬜`imageLight()`

⬜`panorama()`

⬜`lights()`

⬜`lightFalloff()`

⬜`spotLight()`

⬜`noLights()`
</details>

#### Material
<details>
<summary>✴️ In progress (14/15)</summary>

✅`loadShader()`

✅`createShader()`

✅`createFilterShader()`

✅`shader()`

✅`resetShader()`

✅`texture()`

✅`textureMode()`

✅`textureWrap()`

✅`normalMaterial()`

✅`ambientMaterial()`

✅`emissiveMaterial()`

✅`specularMaterial()`

✅`shininess()`

✅`metalness()`

⬜`p5.Shader`
</details>

#### Camera
<details>
<summary>⬜ In progress (0/8)</summary>

⬜`camera()`

⬜`perspective()`

⬜`linePerspective()`

⬜`ortho()`

⬜`frustum()`

⬜`createCamera()`

⬜`p5.Camera`

⬜`setCamera()`
</details>









