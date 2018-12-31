PATH=$PATH:$PWD
TEST_COUNT=0
TEST_PASS=0

TEST_COUNT=$((TEST_COUNT + 1))
if metric-tag-enrichment < testing/data/example-event.json; then
  echo "PASS: succesfully processed an event containing no metrics points (no-op)."
  echo ""
  TEST_PASS=$((TEST_PASS + 1))
else
  echo "FAIL: could not process an event containing no metric points."
  echo ""
fi

TEST_COUNT=$((TEST_COUNT + 1))
if metric-tag-enrichment < testing/data/example-event-with-extracted-metrics.json; then
  echo "PASS: succesfully processed an event containing metrics points."
  echo ""
  TEST_PASS=$((TEST_PASS + 1))
else
  echo "FAIL: could not process an event containing metric points."
  echo ""
fi

if [[ $TEST_COUNT -eq $TEST_PASS ]]; then
  echo "PASS: successfully completed $TEST_PASS of $TEST_COUNT tests."
  echo ""
else
  echo "FAIL: completed $TEST_PASS of $TEST_COUNT tests."
  echo ""
fi
