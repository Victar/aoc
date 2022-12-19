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

	public static final int MINUTES_SILVER = 24;
	public static final int MINUTES_GOLD = 32;

	static long DP_hit = 0;
	int currentMax = Integer.MIN_VALUE;
	State currentMaxState = null;
	Set<String> DP_SET = new HashSet<>();

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		runAny(false);
	}

	@Test public void runGold() throws Exception {
		runAny(true);
	}

	public void runAny(boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		List<Blueprint> blueprintList = new ArrayList<>();
		for (final String input : data) {
			Blueprint blueprint = new Blueprint(input);
			blueprintList.add(blueprint);
		}
		int silverAnwer = 0;
		int goldAnswer = 1;
		int count = isGold ? Math.min(blueprintList.size(), 3) : blueprintList.size();

		for (int i = 0; i < count; i++) {
			System.out.println("##### BLUEPRINT: " + i + " ##### ");
			DP_SET.clear();
			currentMax = -1;
			currentMaxState = null;
			Blueprint currentBlueprint = blueprintList.get(i);
			State startState = new State(currentBlueprint);
			startGame(startState, isGold ? MINUTES_GOLD : MINUTES_SILVER);
			System.out.println("max: " + currentMax + "  " + currentMaxState);
			goldAnswer = goldAnswer * currentMax;
			silverAnwer = silverAnwer + currentMax * currentMaxState.blueprint.getId();
		}
		System.out.println(isGold ? goldAnswer : silverAnwer);
	}

	public boolean currentStateCanWin(final State state, int totalMinutes) {
		int geoToWin = currentMax - state.geodeTotal;
		return geoToWin <= maxGeoInDays(state, totalMinutes);
	}

	public int maxGeoInDays(State state, int totalMinutes) {
		int minutesLeft = state.minutesLeft(totalMinutes);
		int currentGeodeRobotCount = state.geodeRobotCount;
		int result = currentGeodeRobotCount;
		for (int i = 0; i <= minutesLeft; i++) {
			result = result + i + currentGeodeRobotCount;
		}
		return result;
	}

	public void startGame(State current, int totalMinutes) {
		if (current.getMinute() >= totalMinutes) {
			if (current.geodeTotal > currentMax) {
				currentMax = Math.max(currentMax, current.geodeTotal);
				currentMaxState = current;
			}
			return;
		}
		if (!currentStateCanWin(current, totalMinutes)) {
			return;
		}
		List<State> newStatesList = current.getStates();
		for (State state : newStatesList) {

			String id = state.getId();
			if (DP_SET.contains(id)) {
				DP_hit++;
			} else {
				DP_SET.add(id);
				startGame(state, totalMinutes);
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
		State parent = null;

		public State(Blueprint blueprint) {
			this.minute = 0;
			this.oreRobotCount = 1;
			this.blueprint = blueprint;
		}

		public State(Blueprint blueprint, int minute, int oreRobotCount, int clayRobotCount, int obsidianRobotCount, int geodeRobotCount,
		             int oreTotal, int clayTotal, int obsidianTotal, int geodeTotal, State parent) {
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
			this.parent = parent;
		}

		public int minutesLeft(int totalMinutes) {
			return totalMinutes - minute;
		}

		public boolean canBuildGeo() {
			return blueprint.geodeRobotCostOre <= oreTotal && blueprint.geodeRobotCostObsidian <= obsidianTotal;
		}

		public List<State> getStates() {
			ArrayList<State> states = new ArrayList<>();
			int nextMinute = minute + 1;
			int totalOre = oreTotal + oreRobotCount;
			int totalClay = clayTotal + clayRobotCount;
			int totalObsidian = obsidianTotal + obsidianRobotCount;
			int totalGeodeTotal = geodeTotal + geodeRobotCount;
			if (canBuildGeo()) { // Build geo if can
				State state = new State(blueprint, nextMinute, //
						oreRobotCount, clayRobotCount, obsidianRobotCount, geodeRobotCount + 1, // robots
						totalOre - blueprint.geodeRobotCostOre, totalClay, totalObsidian - blueprint.geodeRobotCostObsidian, //matirials
						totalGeodeTotal, this);
				states.add(state);
				return states;
			}

			int MAX_SIZE = 2;
			for (int rO = 0; rO < MAX_SIZE; rO++) {
				for (int rC = 0; rC < MAX_SIZE; rC++) {
					for (int rB = 0; rB < MAX_SIZE; rB++) {
						int totalOreRobots =
								rO * blueprint.oreRobotCostOre + rC * blueprint.clayRobotCostOre + rB * blueprint.obsidianRobotCostOre;
						int totalClayRobots = rB * blueprint.obsidianRobotCostClay;
						int totalRobotsPerMinute = rO + rC + rB;
						if (totalOreRobots <= this.oreTotal && totalClayRobots <= this.clayTotal) {
							if (totalRobotsPerMinute == 1) {  //build Ore, Clay, Obsidian robots if can
								State state = new State(blueprint, nextMinute, //
										oreRobotCount + rO, clayRobotCount + rC, obsidianRobotCount + rB, geodeRobotCount, // robots
										totalOre - totalOreRobots, totalClay - totalClayRobots, totalObsidian, totalGeodeTotal, this);
								states.add(state);
							}
							if (totalRobotsPerMinute == 0 && totalOre
									<= blueprint.maxOre) { //to reduce search add state without building robots only if not enough ore
								State state = new State(blueprint, nextMinute, //
										oreRobotCount + rO, clayRobotCount + rC, obsidianRobotCount + rB, geodeRobotCount,
										// robots rO=rC=rB=0
										totalOre - totalOreRobots, totalClay - totalClayRobots, totalObsidian, totalGeodeTotal, this);
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
		int maxOre;

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
			maxOre = oreRobotCostOre + clayRobotCostOre + obsidianRobotCostOre + geodeRobotCostOre;
		}
	}

}
