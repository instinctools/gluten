export class Constraints {
  public static get baseURL(): string {
    return "http://localhost:8080/api/";
  }

  public static get executionURL(): string {
    return "executions/"
  }

  public static get stopExecution(): string {
    return "stop/"
  }

  public static get nodes(): string {
    return "nodes/"
  }

  //temp
  public static get aggregateStatistic(): string {
    return "aggstatistic?type=request"
  }
}
