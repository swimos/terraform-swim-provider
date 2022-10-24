package swim.basic;

import swim.api.SwimLane;
import swim.api.agent.AbstractAgent;
import swim.api.lane.MapLane;
import swim.api.lane.ValueLane;
import swim.recon.Recon;
import swim.structure.Value;

public class UnitAgent extends AbstractAgent {

  @SwimLane("state")
  ValueLane<Value> state = this.<Value>valueLane()
      .didSet((newValue, oldValue) -> {
        if(!oldValue.isDefinite()){
          logMessage("Container id set to " + Recon.toString(newValue));
        }
        else
        {
          if (newValue.isDefinite()) {
            logMessage("Container id changed from " + Recon.toString(oldValue) + " to " + Recon.toString(newValue));
          }
          else {
            logMessage("Container id cleared");
          }
        }
      });

  @SwimLane("cart")
  MapLane<String, String> cart = this.<String, String>mapLane()
       .didUpdate((key, newValue, oldValue) -> {
         logMessage("Item " + key + " value changed to " + newValue + " from " + oldValue);
       })
       .didRemove((key, oldValue) -> {
         logMessage("Item removed <" + key + "," + oldValue + ">");
       });

  @SwimLane("state2")
  ValueLane<Value> state2 = this.<Value>valueLane()
       .didSet((newValue, oldValue) -> {
         if(!oldValue.isDefinite()){
           logMessage("Container id2 set to " + Recon.toString(newValue));
         }
         else
         {
           if (newValue.isDefinite()) {
             logMessage("Container id2 changed from " + Recon.toString(oldValue) + " to " + Recon.toString(newValue));
           }
           else {
             logMessage("Container id2 cleared");
           }
         }
       });

  @Override
  public void didStart() {
    logMessage("did start");
    //Insert some dummy values into the state of the web agent
  }

  private void logMessage(final String message) {
    System.out.println("/agent" + ": " + message);
  }

}
