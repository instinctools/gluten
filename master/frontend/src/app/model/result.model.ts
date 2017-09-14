import {Metric} from "./metric.model";
export class Result {

  public Id: string;
  public Created: number;
  public Metrics: Metric[];
  public ExecutionID: string;
  public StepType: string;
}
