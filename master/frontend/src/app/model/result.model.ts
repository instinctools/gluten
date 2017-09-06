import {Metric} from "./metric.model";
export class Result {

  public Status: string;
  public Metrics: Metric[];
  public ExecutionID: string;
  public StepType: string;
}
