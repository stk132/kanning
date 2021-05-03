# kanning

## これは何？

mackerel-agent-pluginが出力する

```
メトリクス名\t値\tタイムスタンプ
```

というフォーマットのデータを標準入力から与えると、1つのjsonオブジェクトに変換して出力するツール

## 何のために使うのか？

telegrafの入力用メトリクスをmackerel-agent-pluginから取得したいため、一度jsonに変換する必要があった。

## サンプル

入力データ

```text:_test/fireworq_metrics.tsv
fireworq.jobs.elapsed.jobs_average_elapsed_time 0       1620055407
fireworq.node.active_nodes      1       1620055407
fireworq.node.active_nodes_percentage   100     1620055407
fireworq.queue.workers.queue_idle_workers       20      1620055407
fireworq.queue.workers.queue_running_workers    0       1620055407
fireworq.queue.buffer.queue_outstanding_jobs    0       1620055407
fireworq.jobs.jobs_failure      0       1620055407
fireworq.jobs.jobs_success      0       1620055407
fireworq.jobs.jobs_outstanding  0       1620055407
fireworq.jobs.jobs_waiting      0       1620055407
fireworq.jobs.events.jobs_events_pushed 0       1620055407
fireworq.jobs.events.jobs_events_popped 0       1620055407
fireworq.jobs.events.jobs_events_failed 0       1620055407
fireworq.jobs.events.jobs_events_succeeded      0       1620055407
fireworq.jobs.events.jobs_events_completed      0       1620055407
```

実行コマンド

```shell
cat _test/fireworq_metrics.tsv | kanning | jq .
```

出力

```json
{
  "jobs": {
    "elapsed": {
      "jobs_average_elapsed_time": 0
    },
    "events": {
      "jobs_events_completed": 0,
      "jobs_events_failed": 0,
      "jobs_events_popped": 0,
      "jobs_events_pushed": 0,
      "jobs_events_succeeded": 0
    },
    "jobs_failure": 0,
    "jobs_outstanding": 0,
    "jobs_success": 0,
    "jobs_waiting": 0
  },
  "node": {
    "active_nodes": 1,
    "active_nodes_percentage": 100
  },
  "queue": {
    "buffer": {
      "queue_outstanding_jobs": 0
    },
    "workers": {
      "queue_idle_workers": 20,
      "queue_running_workers": 0
    }
  }
}

```