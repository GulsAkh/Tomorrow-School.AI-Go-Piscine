if [ -z "$HERO_ID" ]; then
    break;
else
    curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r --arg id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives'
fi