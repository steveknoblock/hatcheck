#!/usr/bin/env bash
# hatcheck.sh: start, stop, restart and check the Hatcheck server.
# Keep this script in the Hatcheck project directory (next to hatcheck-server).
#
#   ./hatcheck.sh build     compile ./server into ./hatcheck-server
#   ./hatcheck.sh start     run it in the background (builds first if missing)
#   ./hatcheck.sh stop      stop it (SIGTERM, then wait)
#   ./hatcheck.sh restart
#   ./hatcheck.sh status
#   ./hatcheck.sh logs      follow the log (Ctrl-C to leave)
#
# Settings (HATCHECK_DATA, HATCHECK_META, HATCHECK_UI) are read from .env if present.
set -u

cd "$(dirname "$(readlink -f "$0")")" || exit 1
bin=./hatcheck-server
pidfile=.hatcheck.pid
logfile=hatcheck.log
name=hatcheck-server

# PID of the running server, or nothing. Uses the pid file, falling back to
# a process-name search so a server started by hand is still found.
running_pid() {
  local pid
  if [[ -f $pidfile ]]; then
    pid=$(<"$pidfile")
    if [[ -n $pid ]] && [[ $(ps -p "$pid" -o comm= 2>/dev/null) == "$name" ]]; then
      echo "$pid"; return 0
    fi
  fi
  pgrep -x "$name" | head -n1
}

do_build() {
  echo "Building $bin ..."
  go build -o "$bin" ./server/
}

do_start() {
  local pid; pid=$(running_pid)
  if [[ -n $pid ]]; then echo "Already running (pid $pid)."; return 0; fi
  [[ -x $bin ]] || do_build || return 1
  if [[ -f .env ]]; then set -a; source .env; set +a; fi
  setsid nohup "$bin" >>"$logfile" 2>&1 < /dev/null &
  echo $! > "$pidfile"
  sleep 1
  pid=$(running_pid)
  if [[ -n $pid ]]; then
    echo "Started (pid $pid). Log: $PWD/$logfile"
    echo "UI: http://localhost:8090/ui/index.html"
  else
    echo "Failed to start. Last log lines:" >&2
    tail -n 10 "$logfile" >&2
    rm -f "$pidfile"
    return 1
  fi
}

do_stop() {
  local pid; pid=$(running_pid)
  if [[ -z $pid ]]; then echo "Not running."; rm -f "$pidfile"; return 0; fi
  kill "$pid"
  for _ in $(seq 1 50); do          # wait up to 10 seconds
    kill -0 "$pid" 2>/dev/null || { echo "Stopped."; rm -f "$pidfile"; return 0; }
    sleep 0.2
  done
  echo "Still running after 10s (pid $pid). Force it with: kill -9 $pid" >&2
  return 1
}

do_status() {
  local pid; pid=$(running_pid)
  if [[ -n $pid ]]; then echo "Running (pid $pid)."; else echo "Not running."; return 1; fi
}

case "${1:-}" in
  build)   do_build ;;
  start)   do_start ;;
  stop)    do_stop ;;
  restart) do_stop && do_start ;;
  status)  do_status ;;
  logs)    tail -n 50 -f "$logfile" ;;
  *)       echo "Usage: $0 {build|start|stop|restart|status|logs}" >&2; exit 2 ;;
esac
