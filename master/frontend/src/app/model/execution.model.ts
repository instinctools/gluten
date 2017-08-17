import {ExecutionResult} from "./execution-result.model";
/**
 * Created by prokop06 on 17.8.17.
 */
export class Execution {

  public ID: number;
  public Created: number;
  public Parameters: string;
  public Result: ExecutionResult;
  public ResultID: number;
}
