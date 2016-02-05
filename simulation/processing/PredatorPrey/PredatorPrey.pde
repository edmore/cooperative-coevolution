// Prey Capture simulation
// Written by Edmore Moyo
// Credit Jason Peck for gridlines.pde : https://gist.github.com/gatesphere/4486700

// grid specific vars
// grid spacing in pixels
int GRID_MINOR_SPACING = 10;
int GRID_MAJOR_SPACING = 100;
// gridline colors
color GRID_MINOR_COLOR = color(147, 161, 247, 127);
color GRID_MAJOR_COLOR = color(0, 19, 137, 127);
// gridline weights in pixels
int GRID_MINOR_WEIGHT = 1;
int GRID_MAJOR_WEIGHT = 2;
// when true, grid is drawn, else grid is not
boolean GRID_DRAW = true;
// sketch vars go here
int background_color = 255; //white

// Simulation Variables : Declare some global variables
JSONArray states;
JSONObject state;
float PreyX; // x position of the prey
float PreyY; // y position of the prey
int counter;

void setup() {
  // Load the JSON file as a JSON object
  states = loadJSONArray("../../processingjs/json/states.json");
  size(1000,1000, P2D);
  smooth();
  frameRate(30);
  counter = 0;
}

void draw()
{ 
  background(255, 204, 0);
  // get and set the new state
  if (counter < states.size()){
     state = states.getJSONObject(counter);
     println(state);
  }else{
    noLoop();
  }
  
  drawGrid();
  drawPrey();
  //drawPredator();
  
  counter++;
  delay(10);
}

void drawPrey(){
  fill(0.2);
  noStroke();
  rectMode(CORNER);
  rect(0, 0, 10, 10);
}

// grid code
public void drawGrid() {
  if (GRID_DRAW) {
    int num_minor_x = width/GRID_MINOR_SPACING;
    int num_minor_y = height/GRID_MINOR_SPACING;
    int num_major_x = width/GRID_MAJOR_SPACING;
    int num_major_y = height/GRID_MAJOR_SPACING;
    pushStyle();
    strokeCap(PROJECT);
    strokeWeight(GRID_MINOR_WEIGHT);
    stroke(GRID_MINOR_COLOR);
    for (int i = 0; i < num_minor_y; i++) {
      int y = i * GRID_MINOR_SPACING;
      line(0, y, width, y);
    }
    for (int i = 0; i < num_minor_x; i++) {
      int x = i * GRID_MINOR_SPACING;
      line(x, 0, x, height);
    }
    strokeWeight(GRID_MAJOR_WEIGHT);
    stroke(GRID_MAJOR_COLOR);
    for (int i = 0; i < num_major_y; i++) {
      int y = i * GRID_MAJOR_SPACING;
      line(0, y, width, y);
    }
    for (int i = 0; i < num_major_x; i++) {
      int x = i * GRID_MAJOR_SPACING;
      line(x, 0, x, height);
    }
    popStyle();
  }
}
