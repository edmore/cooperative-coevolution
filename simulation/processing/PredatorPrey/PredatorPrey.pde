// Prey-Capture simulation
// Written by Edmore Moyo

// Declare some global variables
JSONArray states;
JSONObject state;
int PreyX; // x position of the prey
int PreyY; // y position of the prey
JSONArray PredatorX; // position X of the pradators
JSONArray PredatorY; // position Y of the predators

int counter;

void setup() {
  // Load the JSON file as a JSON object
  states = loadJSONArray("../../processingjs/json/states.json");
  size(200,200, P2D); // toroidal grid world
  smooth();
  frameRate(30);
  counter = 0;
}

void draw()
{ 
  background(#0080ff);
  // get and set the new state
  if (counter < states.size()){
     state = states.getJSONObject(counter);
     println(state);
  }else{
    noLoop();
  }
  
  // position of prey
  PreyX = state.getInt("PreyX");
  PreyY = state.getInt("PreyY");

  // position(s) of the predator(s)
  PredatorX = state.getJSONArray("PredatorX");
  PredatorY = state.getJSONArray("PredatorY");

  drawPrey();
  drawPredator();
  
  counter++;
  delay(10);
}

void drawPrey(){
  fill(#ff0000);
  rect(PreyX * 2, PreyY * 2 , 10, 10);
}

void drawPredator(){
  for (int i=0; i < PredatorX.size(); i++){
    fill(#FFFFFF);
    rectMode(CORNER);
    rect(PredatorX.getInt(i) * 2, PredatorY.getInt(i) * 2, 10, 10);
  }
}
