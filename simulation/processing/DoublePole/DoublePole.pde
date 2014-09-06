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
  size(480, 200, P2D);
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
     //println(state);
  }else{
    noLoop();
  }
  X = state.getFloat("X") * 100; 
  Theta1 = state.getFloat("Theta1"); 
  Theta2 = state.getFloat("Theta2"); 
  drawCentreLine();
  drawCart();
  drawPoles();
  
  println("{X : " + X + " (cm), " + "Theta1 : " + Theta1 + ", "  + "Theta2 : " + Theta2 + "}");
  counter++;
  delay(200);
}

void drawCentreLine()
{
  strokeWeight(1);
  line(240, 0, 240, 200);
}

void drawPoles()
{
  stroke(0, 102, 0);
  strokeWeight(4);
  smooth();
  // short pole
  pushMatrix();
  // move the origin to the pivot point
  translate(230+X, 150);
  rectMode(CORNER);
  // then pivot the grid - as this is what moves
  rotate(Theta2 + radians(180));
  fill(0.2);
  // and draw the square at the new origin
  rect(0, 0, 1, 10);
  popMatrix();

  // long pole
  pushMatrix();
  // move the origin to the pivot point
  translate(250+X, 150);
  rectMode(CORNER);
  // then pivot the grid - as this is what moves
  rotate(Theta1 + radians(180));
  fill(153);
  // and draw the square at the new origin
  rect(0, 0, 1, 100);
  popMatrix();
}

void drawCart(){
  fill(0.2);
  noStroke();
  rectMode(CORNER);
  rect(190+X, 150, 100, 50);
}
