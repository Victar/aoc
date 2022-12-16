package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day16 extends BaseTest {

	public static final int DAY = 16;
	public static int skipGold = 0;
	static int MAX_KNOWN = 0;
	static Map<String, Valve> mapValve = new HashMap<>();
	static int totalPresure = 0;
	static long DP_hit = 0;
	Set<String> DP_SET = new HashSet<>();
	Map<String, Integer> DP_MAP = new HashMap<>();

	int currentMax = Integer.MIN_VALUE;
	State currentMaxState;
	Map<Integer, Integer> borders = new HashMap<>();

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runBoth() throws Exception {
		runAny(false);
		currentMaxState.populateBoarder(borders);
		DP_SET.clear();
		currentMax = Integer.MIN_VALUE;
		currentMaxState = null;
		DP_hit = 0;
		DP_MAP.clear();
		MAX_KNOWN = 2905;
		totalPresure = 0;
		runAny(true);
		System.out.println("#####");
		//		2371 - skip Gold 10
		//		2318 - skip Gold 12
		//      2787 - skip Gold 8 270 000 000
		//      2787 - skip Gold 6 350 000 000
		//      2865 - skip Gold 5 xxx 000 000
		//      2905 - skip Gold 4 xxx 000 000
		//      2905 - skip Gold 3  80 000 000
		//      2905 - skip Gold 2  80 000 000
		// 2911

		System.out.println(skipGold * currentMaxState.releasingPressure + currentMaxState.totalReleasedPressure);

	}

	public void runAny(boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/sample.txt");
		for (final String input : data) {
			Valve.initValve(mapValve, input);
		}
		State start = new State(1, mapValve.get("AA"), mapValve.get("AA"), 0, 0, new HashSet<>(), null);
		startGame(start, isGold);
		System.out.println(currentMax);
		currentMaxState.print();

	}

	public boolean borderCheckPass(State current) {
		if (borders.isEmpty()) {
			return true;
		}
		Integer borderLimit = borders.get(current.minute);
		return borderLimit == null || borderLimit <= current.totalReleasedPressure;
	}

	public boolean bestCheckPass(State current) {
		String idForBestCheck = current.getIdForBestCheck();
		int currentValue = current.getTotalReleasedPressure();

		if (currentValue + totalPresure * (26 - current.minute) < MAX_KNOWN) {
			if (DP_hit % 100000 == 0) {
				System.out.println("bestCheckPass not passed for : " + current);
			}
			return false;
		}
		if (DP_MAP.containsKey(idForBestCheck)) {
			int bestValue = DP_MAP.get(idForBestCheck);
			if (currentValue > bestValue) {
				DP_MAP.put(idForBestCheck, currentValue);
				return true;
			} else {
				return false;
			}
		} else {
			DP_MAP.put(idForBestCheck, currentValue);
		}
		return true;
	}

	public void startGame(State current, boolean isGold) {
		if (current.getMinute() >= (isGold ? 26 - skipGold : 30)) {
			if (current.getTotalReleasedPressure() > currentMax) {
				currentMax = Math.max(currentMax, current.getTotalReleasedPressure());
				currentMaxState = current;
				//				System.out.println("NEW max");
				//				System.out.println(currentMax);
				//				currentMaxState.print();
			}
			return;
		}
		if (!borderCheckPass(current)) {
			return;
		}

		if (!(bestCheckPass(current))) {
			return;
		}

		List<State> newStatesList = isGold ? current.getNewStateGold() : current.getNewState(true);

		for (State state : newStatesList) {
			String id = state.getId();
			if (DP_SET.contains(id)) {
				DP_hit++;
				if (DP_hit % 10000000 == 0) {
					System.out.println(DP_hit);
				}
			} else {
				DP_SET.add(id);
				startGame(state, isGold);
			}
		}
	}

	@Data static class State {

		int minute;
		int releasingPressure;
		int totalReleasedPressure;
		Set<String> openedValveIds;
		Valve currentValve;
		Valve currentValveElefant;

		String openedValveId;

		State parent = null;

		public State(int minute, Valve currentValve, Valve currentValveElefant, int releasingPressure, int totalReleasedPressure,
		             Set<String> openedValveIds, State parent) {
			this.minute = minute;
			this.currentValve = currentValve;
			this.currentValveElefant = currentValveElefant;
			this.releasingPressure = releasingPressure;
			this.totalReleasedPressure = totalReleasedPressure;
			this.openedValveIds = openedValveIds;
			this.parent = parent;
		}

		public List<State> getNewStateGold() {
			List<State> stateList = new ArrayList<>();

			List<State> statesFirstPlayer = getNewState(true);
			for (State stateFirstPlayer : statesFirstPlayer) {
				stateList.addAll(stateFirstPlayer.getNewState(false));
			}
			return stateList;
		}

		public List<State> getNewState(boolean firstPlayer) {
			int newMinute = firstPlayer ? this.minute + 1 : this.minute;
			Valve currentActor = firstPlayer ? this.currentValve : this.currentValveElefant;
			Valve secondValve = firstPlayer ? this.currentValveElefant : this.currentValve;
			State parent = firstPlayer ? this : this.parent;
			int newReleasingPressure = firstPlayer ? this.releasingPressure : 0;
			Set<String> newList = new HashSet<>(openedValveIds);
			List<State> stateList = new ArrayList<>();

			State newState;
			if (firstPlayer) {
				newState = new State(newMinute, currentActor, secondValve, releasingPressure, totalReleasedPressure + newReleasingPressure,
						newList, parent);
				stateList.add(newState);

			} else {
				newState = new State(newMinute, secondValve, currentActor, releasingPressure, totalReleasedPressure + newReleasingPressure,
						newList, parent);
				stateList.add(newState);

			}
			if (totalPresure == releasingPressure) {
				return stateList;
			}
			if (!this.openedValveIds.contains(currentActor.getName())) {
				newList = new HashSet<>(openedValveIds);
				newList.add(currentActor.name);
				if (firstPlayer) {
					newState = new State(newMinute, currentActor, secondValve, releasingPressure + currentActor.getRate(),
							totalReleasedPressure + newReleasingPressure + currentActor.getRate(), newList, parent);
					stateList.add(newState);

				} else {
					newState = new State(newMinute, secondValve, currentActor, releasingPressure + currentActor.getRate(),
							totalReleasedPressure + newReleasingPressure + currentActor.getRate(), newList, parent);
					stateList.add(newState);
				}
			}
			for (String newPath : currentActor.paths) {
				newList = new HashSet<>(openedValveIds);
				Valve newValve = mapValve.get(newPath);
				if (firstPlayer) {
					newState = new State(newMinute, newValve, secondValve, releasingPressure, totalReleasedPressure + newReleasingPressure,
							newList, parent);
					stateList.add(newState);
				} else {
					newState = new State(newMinute, secondValve, newValve, releasingPressure, totalReleasedPressure + newReleasingPressure,
							newList, parent);
					stateList.add(newState);
				}
			}

			//			if (DP_hit % 1000000 == 0) {
			//				System.out.println("new List size: " + newList.size() + " stateList: " + stateList.size() + " " + this.getId());
			//			}
			return stateList;
		}

		//
		public String getIdForBestCheck() {
			return "m:" + this.minute + ";" + currentValve.getName() + ";" + currentValveElefant.getName() + ";" + this.releasingPressure;
		}

		public String getId() {

			return "m:" + this.minute + ";" + currentValve.getName() + ";" + currentValveElefant.getName() + ";" + this.releasingPressure
					+ ";" + this.totalReleasedPressure + ";"; //Gold
		}

		public void print() {

			System.out.println(this.getId());
			System.out.println("== Minute " + this.minute + " ==");
			System.out.println("Valve  " + this.openedValveIds + " open, releasing " + this.releasingPressure + " pressure. total: "
					+ this.totalReleasedPressure + ".");
			System.out.println("Current: " + this.currentValve.name + "  Current Elefant: " + this.currentValveElefant.name);
			System.out.println();

			if (this.parent != null) {
				parent.print();
			}
		}

		public void populateBoarder(Map<Integer, Integer> borders) {
			if (!borders.containsKey(this.minute)) {
				System.out.println("Added: " + minute + "  " + this.totalReleasedPressure);
				borders.put(minute, totalReleasedPressure);
			}
			if (this.parent != null) {
				this.parent.populateBoarder(borders);
			}
		}
	}

	@Data static class Valve {

		String name;
		int rate;
		List<String> paths;

		//		Valve JJ has flow rate=21; tunnel leads to valve II
		public static void initValve(Map<String, Valve> mapValve, String currentString) {
			boolean multiple = currentString.contains("valves");
			String name = currentString.substring(6, 8);
			Valve current = mapValve.containsKey(name) ? current = mapValve.get(name) : new Valve();
			current.name = name;
			String[] spl = currentString.split("; tunnel");
			current.rate = Integer.parseInt(spl[0].substring(23));
			totalPresure += current.rate;
			String[] ps = spl[1].substring(multiple ? 17 : 16).split(", ");
			current.paths = Arrays.asList(ps);
			mapValve.put(name, current);
		}
	}

}
