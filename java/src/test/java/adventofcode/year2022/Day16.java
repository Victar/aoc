package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day16 extends BaseTest {

	public static final int DAY = 16;
	static Map<String, Valve> mapValve = new HashMap<>();
//	Map<String, State> DP = new HashMap<>();
	Set<String> DP_SET = new HashSet<>();
	static int totalPresure = 0;

	int currentMax;
	State currentMaxState;

	long DP_hit = 0;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/sample.txt");
		for (final String input : data) {
			//			System.out.println(input);
			Valve.initValve(mapValve, input);
		}
		State start = new State(1, mapValve.get("AA"), 0, 0, new HashSet<>(), null);
		startGame(start);
		System.out.println(currentMax);
		currentMaxState.print();

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

	public void startGame(State current) {
		//		System.out.println(current + " " + DP_hit );
		if (current.getMinute() >= 30) {
			if (current.getTotalReleasedPressure()> currentMax){
				currentMax = Math.max(currentMax, current.getTotalReleasedPressure());
				currentMaxState = current;
			}
			return;
		}
		//Open valves
//		List<State> newStatesList = current.getNewStateGold();
		List<State> newStatesList = current.getNewState();

		for (State state : newStatesList) {
			String id = state.getId();
			if (DP_SET.contains(id)) {
				DP_hit++;
				if (DP_hit % 1000000 ==0){
					System.out.println(DP_hit);
				}
			} else {
				DP_SET.add(id);
				startGame(state);
			}
		}
	}

	@Data static class State {

		int minute;
		Valve currentValve;
		Valve currentValveElefant;

		int releasingPressure;
		int totalReleasedPressure;
		Set<String> openedValveIds;
		List<String> openedValveIdsList;

		State parent = null;

		public State(int minute, Valve currentValve, int releasingPressure, int totalReleasedPressure, Set<String> openedValveIds, State parent) {
			this.minute = minute;
			this.currentValve = currentValve;
			this.releasingPressure = releasingPressure;
			this.totalReleasedPressure = totalReleasedPressure;
			this.openedValveIds = openedValveIds;
			this.parent = parent;
			openedValveIdsList = new ArrayList<>(openedValveIds);
			Collections.sort(openedValveIdsList);

		}

		public State(int minute, Valve currentValve,  Valve currentValveElefant, int releasingPressure, int totalReleasedPressure, Set<String> openedValveIds, State parent) {
			this.minute = minute;
			this.currentValve = currentValve;
			this.releasingPressure = releasingPressure;
			this.totalReleasedPressure = totalReleasedPressure;
			this.openedValveIds = openedValveIds;
			this.currentValveElefant = currentValveElefant;
			this.parent = parent;
			openedValveIdsList = new ArrayList<>(openedValveIds);
			Collections.sort(openedValveIdsList);

		}

		public List<State> getNewStateGold() {
			List<State> stateList = new ArrayList<>();
			Set<String> newList = new HashSet<>(openedValveIds);
			State newState =
					new State(this.minute + 1, this.currentValve, this.currentValveElefant, releasingPressure, totalReleasedPressure + releasingPressure, newList, this);
			stateList.add(newState);
			if (totalPresure == releasingPressure){
				return stateList;
			}
			if (!this.openedValveIds.contains(currentValve.getName())) {
				newList = new HashSet<>(openedValveIds);
				newList.add(currentValve.name);
				newState = new State(this.minute + 1, this.currentValve, releasingPressure + currentValve.getRate(),
						totalReleasedPressure + releasingPressure + currentValve.getRate(), newList, this);
				stateList.add(newState);
			}
			for (String newPath : currentValve.paths) {
				newList = new HashSet<>(openedValveIds);
				Valve newValve = mapValve.get(newPath);
				if (newValve != null) {
					newState =
							new State(this.minute + 1, newValve, releasingPressure, totalReleasedPressure + releasingPressure, newList, this);
					stateList.add(newState);
				} else {
					System.out.println(newValve);
				}
			}

			return stateList;
		}

		public List<State> getNewState() {
			List<State> stateList = new ArrayList<>();
			Set<String> newList = new HashSet<>(openedValveIds);
			State newState =
					new State(this.minute + 1, this.currentValve, releasingPressure, totalReleasedPressure + releasingPressure, newList, this);
			stateList.add(newState);
			if (totalPresure == releasingPressure){
				return stateList;
			}
			if (!this.openedValveIds.contains(currentValve.getName())) {
				newList = new HashSet<>(openedValveIds);
				newList.add(currentValve.name);
				newState = new State(this.minute + 1, this.currentValve, releasingPressure + currentValve.getRate(),
						totalReleasedPressure + releasingPressure + currentValve.getRate(), newList, this);
				stateList.add(newState);
			}
				for (String newPath : currentValve.paths) {
					newList = new HashSet<>(openedValveIds);
					Valve newValve = mapValve.get(newPath);
					if (newValve != null) {
						newState =
								new State(this.minute + 1, newValve, releasingPressure, totalReleasedPressure + releasingPressure, newList, this);
						stateList.add(newState);
					} else {
						System.out.println(newValve);
					}
				}

			return stateList;
		}

		public String getId() {


			return "m:" + this.minute + ";" + currentValve.getName() + ";" + this.releasingPressure + ";" + this.totalReleasedPressure + ";"; //Silver
//			return "m:" + this.minute + ";" + currentValve.getName() + ";" + currentValveElefant.getName() + ";"+ this.releasingPressure + ";" + this.totalReleasedPressure + ";"; //Silver

//			return "m:" + this.minute + ";" + currentValve.getName() + ";" + this.releasingPressure + ";" + this.totalReleasedPressure + ";"
//					+ String.join(",",  openedValveIds);
		}



		public void print(){

			System.out.println(this.getId());
			System.out.println("== Minute "+this.minute +" ==");
			System.out.println("Valve  "+this.openedValveIds +" open, releasing " + this.releasingPressure+" pressure.");
			System.out.println("Current: "+this.currentValve.name);
			System.out.println();


			if (this.parent!=null){
				parent.print();
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
			totalPresure +=current.rate;
			String[] ps = spl[1].substring(multiple ? 17 : 16).split(", ");
			current.paths = Arrays.asList(ps);
			mapValve.put(name, current);
			//			System.out.println(current);
			//
		}
	}

}
