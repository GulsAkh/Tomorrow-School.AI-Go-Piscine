ls -l | sed '1d' ; ls -l | awk 'NR == 1 || (NR > 1 && FNR % 2 == 0)'