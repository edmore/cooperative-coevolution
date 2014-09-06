// Declare some global variables
JSONArray states;
JSONObject state;
float X; // position of cart
float Theta1; // long pole angle
float Theta2; // short pole angle
int counter;

void setup(){
  // Load the JSON file as a JSON object
  states = loadJSONArray("../../processingjs/json/states.json");
  size(480, 200);
  noStroke();
  smooth();
  //noLoop();
  counter = 0;
}

void draw()
{ 
  background(102);
  // get and set the new state
  if (counter < states.size()){
     state = states.getJSONObject(counter);
     //println(state);
  }else{
    noLoop();
  }
  X = state.getFloat("X"); 
  Theta1 = state.getFloat("Theta1"); 
  Theta2 = state.getFloat("Theta2"); 
  
  drawCart();
  drawPoles();
  
  println("{X : " + X + " (cm), " + "Theta1 : " + Theta1 + ", "  + "Theta2 : " + Theta2 + "}");
  counter++;
  delay(500);
}

void drawPoles()
{
  stroke(255);
  strokeWeight(2);
  // short pole
  pushMatrix();
  // move the origin to the pivot point
  translate(240+X, 150);
  rectMode(CORNER);
  // then pivot the grid - as this is what moves
  rotate(Theta2 + radians(180));
  fill(0.2);
  // and draw the square at the new origin - set to longer length just for simulation
  rect(0, 0, 1, 20);
  popMatrix();

  // long pole
  pushMatrix();
  // move the origin to the pivot point
  translate(240+X, 150);
  rectMode(CORNER);
  // then pivot the grid - as this is what moves
  rotate(Theta1 + radians(180));
  fill(0.2);
  // and draw the square at the new origin
  rect(0, 0, 1, 100);
  popMatrix();
}

void drawCart(){
  fill(0.2);
  rectMode(CORNER);
  rect(190+X, 150, 100, 50);
}
