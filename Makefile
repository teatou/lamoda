up:
		docker compose -f docker-compose.yaml up -d $(c)

stop:
		docker compose -f docker-compose.yaml stop $(c)

down:
		docker compose -f docker-compose.yaml down
