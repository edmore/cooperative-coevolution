import pygame

background_colour = (255,255,255)
(width, height) = (480, 200)

screen = pygame.display.set_mode((width, height))
pygame.display.set_caption('Double Pole Balancing Task')
screen.fill(background_colour)


pygame.display.flip()

running = True
while running:
  for event in pygame.event.get():
    if event.type == pygame.QUIT:
      running = False
