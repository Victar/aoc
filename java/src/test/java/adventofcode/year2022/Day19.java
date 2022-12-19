package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Day19 extends BaseTest {

	public static final int DAY = 19;

	public static final int MINUTES = 32;

	static long DP_hit = 0;
//	int silverAnwer;
	int currentMax = Integer.MIN_VALUE;
	State currentMaxState = null;
	Set<String> DP_SET = new HashSet<>();

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/sample.txt");
		List<Blueprint> blueprintList = new ArrayList<>();
		for (final String input : data) {
			Blueprint blueprint = new Blueprint(input);
			blueprintList.add(blueprint);
			System.out.println(input);
			System.out.println(blueprint);

		}
		int silverAnwer =0;

		int count = Math.min(blueprintList.size(), 3);
		for (int i = 0; i <count; i++) {
			System.out.println("##### BLUEPRINT: "+i+" ##### " );
			DP_SET.clear();
			currentMax = -1;
			currentMaxState = null;
			Blueprint currentBlueprint = blueprintList.get(i);
			State startState = new State(currentBlueprint);
			startGame(startState);
			System.out.println(currentMaxState);
			silverAnwer = silverAnwer + currentMax * currentMaxState.blueprint.getId();
		}
		System.out.println(silverAnwer);
	}

	public boolean currentStateCanWin(final State state){
		int minLeft = MINUTES - state.minute;
		int geoToWin = currentMax - state.geodeTotal;
		if (geoToWin > maxGeoInDays(minLeft, state.geodeRobotCount)){
			return false;
		}
		return true;
	}
	public int maxGeoInDays(int minutes, int currentRobots){
		int result = currentRobots;
		for (int i=0; i<=minutes; i++){
			result = result + i + currentRobots;
		}
		return result;
	}
	public void startGame(State current) {
		if (current.getMinute() >= MINUTES) {

			if (current.geodeTotal > currentMax) {

				currentMax = Math.max(currentMax, current.geodeTotal);
				currentMaxState = current;
				System.out.println("New max: " + currentMax + "  " + currentMaxState);
			}
			return;
		}
		if (!currentStateCanWin(current)){
//						if (DP_hit%10000==0){
//							System.out.println(current);
//						}
			return;
		}
		List<State> newStatesList = current.getStates();
		for (State state : newStatesList) {

			String id = state.getId();
			if (DP_SET.contains(id)) {
				DP_hit++;
			} else {
				DP_SET.add(id);
				startGame(state);
			}
		}
	}

	@Data static class State {

		int minute;
		int oreRobotCount;
		int clayRobotCount;
		int obsidianRobotCount;
		int geodeRobotCount;

		int oreTotal;
		int clayTotal;
		int obsidianTotal;
		int geodeTotal;
		Blueprint blueprint;

		public State(Blueprint blueprint) {
			this.minute = 0;
			this.oreRobotCount = 1;
			this.blueprint = blueprint;
		}

		public State(Blueprint blueprint, int minute, int oreRobotCount, int clayRobotCount, int obsidianRobotCount, int geodeRobotCount,
		             int oreTotal, int clayTotal, int obsidianTotal, int geodeTotal) {
			//			if (oreRobotCount>1){
			//				System.out.println(oreRobotCount);
			//			}
			this.blueprint = blueprint;
			this.minute = minute;
			this.oreRobotCount = oreRobotCount;
			this.clayRobotCount = clayRobotCount;
			this.obsidianRobotCount = obsidianRobotCount;
			this.geodeRobotCount = geodeRobotCount;
			this.oreTotal = oreTotal;
			this.clayTotal = clayTotal;
			this.obsidianTotal = obsidianTotal;
			this.geodeTotal = geodeTotal;
		}

		public List<State> getStates() {
			ArrayList<State> states = new ArrayList<>();
			int nextMinute = minute + 1;
			int totalOre = oreTotal + oreRobotCount;
			int totalClay = clayTotal + clayRobotCount;
			int totalObsidian = obsidianTotal + obsidianRobotCount;
			int totalGeodeTotal = geodeTotal + geodeRobotCount;
			//			if (totalClay > 0){
			//				System.out.println("Prduciang clay " + totalClay);
			//			}
			//state without build
			int MAX_SIZE = 2;
			for (int rO = 0; rO < MAX_SIZE; rO++) {
				for (int rC = 0; rC < MAX_SIZE; rC++) {
					for (int rB = 0; rB < MAX_SIZE; rB++) { //rB - obsidian;
						for (int rG = 0; rG < MAX_SIZE; rG++) {
							//try to build such robots;
							int totalOreRobots =
									rO * blueprint.oreRobotCostOre + rC * blueprint.clayRobotCostOre + rB * blueprint.obsidianRobotCostOre
											+ rG * blueprint.geodeRobotCostOre;
							int totalClayRobots = rB * blueprint.obsidianRobotCostClay;
							int totalObsidianRobots = rG * blueprint.geodeRobotCostObsidian;
							int totalRobotsPreMinute = rO + rC + rB + rG;
							if (totalRobotsPreMinute < 2 && totalOreRobots <= this.oreTotal && totalClayRobots <= this.clayTotal
									&& totalObsidianRobots <= this.obsidianTotal) {

								State state = new State(blueprint, nextMinute, //
										oreRobotCount + rO, clayRobotCount + rC, obsidianRobotCount + rB, geodeRobotCount + rG, // robots
										totalOre - totalOreRobots, totalClay - totalClayRobots, totalObsidian - totalObsidianRobots,
										totalGeodeTotal);
								states.add(state);
							}
						}
					}
				}
			}

			return states;
		}

		public String getId() {
			return "m" + minute + ":" + oreRobotCount + ":" + clayRobotCount + ":" + obsidianRobotCount + ":" + geodeRobotCount + ":"
					+ oreTotal + ":" + clayTotal + ":" + obsidianTotal + ":" + geodeTotal;
		}
	}

	@Data static class Blueprint {

		int id;
		int oreRobotCostOre;
		int clayRobotCostOre;
		int obsidianRobotCostOre;
		int obsidianRobotCostClay;
		int geodeRobotCostOre;
		int geodeRobotCostObsidian;

		public Blueprint(String input) {
			//			Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 8 clay. Each geode robot costs 2 ore and 15 obsidian.
			String[] str = input.split(": ");
			id = Integer.parseInt(str[0].substring(10));
			String[] robotsStr = str[1].split("\\.");
			oreRobotCostOre = Integer.parseInt(robotsStr[0].substring(21).split(" ")[0]);
			clayRobotCostOre = Integer.parseInt(robotsStr[1].substring(23).split(" ")[0]);
			String[] obsidianStr = robotsStr[2].split(" ");
			obsidianRobotCostOre = Integer.parseInt(obsidianStr[5]);
			obsidianRobotCostClay = Integer.parseInt(obsidianStr[8]);
			String[] geodeStr = robotsStr[3].split(" ");
			geodeRobotCostOre = Integer.parseInt(geodeStr[5]);
			geodeRobotCostObsidian = Integer.parseInt(geodeStr[8]);
		}
	}

}
