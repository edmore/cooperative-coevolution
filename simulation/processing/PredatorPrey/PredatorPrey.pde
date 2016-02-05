// Prey Capture simulation
// Written by Edmore Moyo

// Declare some global variables
JSONArray states;
JSONObject state;
float PreyX; // x position of the prey
float PreyY; // y position of the prey
int counter;

void setup() {
  // Load the JSON file as a JSON object
  states = loadJSONArray("../../processingjs/json/states.json");
  size(100,100, P2D);
  scale(1, -1);
  translate(0, -height);
  smooth();
  frameRate(30);
  counter = 0;
}

void draw()
{ 
  // get and set the new state
  if (counter < states.size()){
     state = states.getJSONObject(counter);
     println(state);
  }else{
    noLoop();
  }
  
  drawPrey();
  //drawPredator();
  
  counter++;
  delay(10);
}

void drawPrey(){
  fill(0.2);
  noStroke();
  rectMode(CORNER);
  rect(0, 0, 1, 1);
}
