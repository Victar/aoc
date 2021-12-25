package adventofcode.year2021;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;


import static adventofcode.year2021.Day23.GameState.*;
import static org.junit.Assert.assertEquals;

public class Day23 extends BaseTest {

	int minPrice = Integer.MAX_VALUE;
	GameState minState;

	@Test public void singleGoInRight() throws Exception {
		final GameState gs = new GameState();
		gs.DD[1] = GameState.S;
		gs.HALL[10] = GameState.D;
		final Map<String, GameState> DP = new HashMap<>();
		findSolution(gs, DP);
		assertEquals(3000, minPrice);
		assertEquals(DONE, minState);
	}

	@Test public void singleGoInLeft() throws Exception {
		final GameState gs = new GameState();
		gs.DD[1] = GameState.S;
		gs.HALL[0] = GameState.D;
		final Map<String, GameState> DP = new HashMap<>();
		findSolution(gs, DP);
		assertEquals(9000, minPrice);
		assertEquals(DONE, minState);
	}

	@Test public void singleGoOut() throws Exception {
		final GameState gs = new GameState();
		gs.AA[1] = GameState.B;
		gs.HALL[10] = GameState.A;
		gs.BB[1] = GameState.S;

		final Map<String, GameState> DP = new HashMap<>();
		findSolution(gs, DP);
		assertEquals(49, minPrice);
		assertEquals(DONE, minState);
	}

	@Test public void runSilver() throws Exception {
		final GameState gsinput = new GameState(new String[] { C, D }, new String[] { A, B }, new String[] { D, C }, new String[] { B, A });
//		final GameState gsinput = new GameState(new String[] { A, B }, new String[] { D, C }, new String[] { C, B }, new String[] { A, D });
		final Map<String, GameState> DP = new HashMap<>();
		findSolution(gsinput, DP);
		System.out.println(this.minPrice);
		System.out.println("DP size: " + DP.size());
		if (this.minState != null) {
			this.minState.printWithParent();
		}
	}

	@Test public void runGold() throws Exception {
				final  GameState gsinput = new GameState(new String[]{C,D,D,D}, new String[]{ A,B,C, B},new String[]{D,A,B,C},new String[]{B,C,A,A});
		final Map<String, GameState> DP = new HashMap<>();
		findSolution(gsinput, DP);
		System.out.println(this.minPrice);
		System.out.println(DP.size());
		if (this.minState != null) {
			this.minState.printWithParent();
		}
	}

	public void findSolution(final GameState gameState, final Map<String, GameState> DP) {
		if (gameState.isDone()) {
			if (this.minPrice > gameState.getPrice()) {
				this.minPrice = gameState.getPrice();
				this.minState = gameState;
				System.out.println("New min: " + this.minPrice);
			}
			return;
		}
		final String key = gameState.getKey();
		if (DP.containsKey(key)) {
			return;
		}
		final List<GameState> states = gameState.getNextGameStates();
		for (final GameState state : states) {
			findSolution(state, DP);
		}
		DP.put(key, gameState);
	}

	@Data static class GameState {

		public static GameState DONE = new GameState();
		private GameState parent;

		int price;

		public static final String A = "A";
		public static final String B = "B";
		public static final String C = "C";
		public static final String D = "D";
		public static final String S = "."; //space
		public static final String R = "x"; //room
		public static final String SR = S + R; //room

		String[] AA = { A, A };
		String[] BB = { B, B };
		String[] CC = { C, C };
		String[] DD = { D, D };
		String[] HALL = { S, S, R, S, R, S, R, S, R, S, S };

		static final Map<String, Integer> DOORS = new HashMap<>();
		static final Map<String, Integer> PRICE = new HashMap<>();

		static {
			DOORS.put(A, 2);
			DOORS.put(B, 4);
			DOORS.put(C, 6);
			DOORS.put(D, 8);
			PRICE.put(A, 1);
			PRICE.put(B, 10);
			PRICE.put(C, 100);
			PRICE.put(D, 1000);
		}

		public GameState(final String[] AA, final String[] BB, final String[] CC, final String[] DD) {
			this.AA = AA;
			this.BB = BB;
			this.CC = CC;
			this.DD = DD;
		}

		public GameState() {
		}

		private String[] getSetByKey(final String key) {
			if (A.equals(key)) {
				return this.AA;
			}
			if (B.equals(key)) {
				return this.BB;
			}
			if (C.equals(key)) {
				return this.CC;
			}
			if (D.equals(key)) {
				return this.DD;
			}
			return null;
		}

		private void moveToCell(final int position) {
			final String current = this.HALL[position];
			int distance = Math.abs(position - DOORS.get(current)) + 1;
			final String[] arr = getSetByKey(current);
			int arrayDistance = arr.length; //2
			boolean found = false;
			for (int i = 0; i < arr.length && !found; i++) {
				if (S.equals(arr[i])) {
					arr[i] = current;
					found = true;
				}
				--arrayDistance;
			}
			this.HALL[position] = S;
			this.price = this.price + PRICE.get(current) * (distance + arrayDistance);

		}

		private boolean canMoveToCell(final String current, final int position) {
			if (A.equals(current) && containsOnlyOrEmpty(this.AA, A)) {
				return directionFree(position, A);
			}
			if (B.equals(current) && containsOnlyOrEmpty(this.BB, B)) {
				return directionFree(position, B);
			}
			if (C.equals(current) && containsOnlyOrEmpty(this.CC, C)) {
				return directionFree(position, C);
			}
			if (D.equals(current) && containsOnlyOrEmpty(this.DD, D)) {
				return directionFree(position, D);
			}
			return false;
		}

		private boolean directionFree(final int start, final String name) {
			boolean result = true;
			final int end = DOORS.get(name);
			if (start < end) {
				for (int i = start + 1; i <= end; i++) {
					if (!StringUtils.containsOnly(this.HALL[i], SR)) {
						result = false;
					}
				}
			} else {
				for (int i = start - 1; i >= end; i--) {
					if (!StringUtils.containsOnly(this.HALL[i], SR)) {
						result = false;
					}
				}
			}
			return result;
		}

		public List<GameState> getNextGameStates() {
			final List<GameState> states = new ArrayList<GameState>();
			if (this.isDone()) {
				return states;
			}
			for (int i = 0; i < this.HALL.length; i++) {
				if (canMoveToCell(this.HALL[i], i)) {
					final GameState gs = copy();
					gs.moveToCell(i);
					states.add(gs);
				}
			}

			states.addAll(goOut(AA, A));
			states.addAll(goOut(DD, D));
			states.addAll(goOut(CC, C));
			states.addAll(goOut(BB, B));
			return states;
		}

		private List<GameState> goOut(final String[] arr, final String type) {
			final List<GameState> states = new ArrayList<>();
			//Check if it contains letter except needed
			boolean needToGoOut = false;
			for (int i = 0; i < arr.length; i++) {
				if (!StringUtils.containsOnly(arr[i], type + S)) {
					needToGoOut = true;
				}
			}
			if (!needToGoOut) {
				return states;
			}
			// Need to go out check distance for first letter to go out
			final int positionArr = DOORS.get(type);
			int distance = 0;
			String letterToGoOut = StringUtils.EMPTY;
			boolean found = false;
			for (int i = arr.length - 1; i >= 0 && !found; i--) {
				letterToGoOut = arr[i];
				distance++;
				if (!S.equals(letterToGoOut)) {
					found = true;
				}
			}
			//			System.out.println(Arrays.toString(arr) + " go out: " + letterToGoOut + " position: " + distance);

			//Go out left
			int distanceInside = 0;
			for (int i = positionArr; i < HALL.length; i++) {
				if (canMoveInHallFromTo(positionArr, i)) {
					final GameState gs = copy();
					gs.setArrayValue(type, distance, S);
					gs.HALL[i] = letterToGoOut;
					gs.price = gs.price + (distance + distanceInside) * PRICE.get(letterToGoOut);
					states.add(gs);
				}
				distanceInside++;

			}
			//Go out right
			distanceInside = 0;
			for (int i = positionArr; i >= 0; i--) {
				if (canMoveInHallFromTo(positionArr, i)) {
					final GameState gs = copy();
					gs.setArrayValue(type, distance, S);
					gs.HALL[i] = letterToGoOut;
					gs.price = gs.price + (distance + distanceInside) * PRICE.get(letterToGoOut);
					states.add(gs);
				}
				distanceInside++;
			}
			return states;
		}

		private void setArrayValue(final String type, final int position, final String value) {
			final String arr[] = getSetByKey(type);
			arr[arr.length - position] = value;
		}

		private boolean canMoveInHallFromTo(final int start, final int end) {
			if (start < end) {
				for (int i = start; i <= end; i++) {
					if (!StringUtils.containsOnly(this.HALL[i], SR)) {
						return false;
					}
				}
				return S.equals(HALL[end]);
			} else {
				for (int i = start; i >= end; i--) {
					if (!StringUtils.containsOnly(this.HALL[i], SR)) {
						return false;
					}
				}
				return S.equals(HALL[end]);
			}
		}

		public GameState copy() {
			final GameState newState = new GameState();
			newState.AA = Arrays.stream(this.AA).toArray(String[]::new);
			newState.BB = Arrays.stream(this.BB).toArray(String[]::new);
			newState.CC = Arrays.stream(this.CC).toArray(String[]::new);
			newState.DD = Arrays.stream(this.DD).toArray(String[]::new);
			newState.HALL = Arrays.stream(this.HALL).toArray(String[]::new);
			newState.price = this.price;
			newState.parent = this;
			return newState;

		}

		public String getKey() {
			return "A" + Arrays.toString(this.AA) + ",B" + Arrays.toString(this.BB) + ",C" + Arrays.toString(this.CC) + ",D"
					+ Arrays.toString(this.DD) + ",H" + Arrays.toString(this.HALL) + "," + this.price;
		}

		public void printWithParent() {
			if (this.parent != null) {
				this.parent.printWithParent();
			}
			print();

		}

		public void print() {
			final StringBuilder sb = new StringBuilder();
			sb.append("#############");
			sb.append(System.lineSeparator());
			sb.append("#");
			for (int i = 0; i < this.HALL.length; i++) {
				sb.append(this.HALL[i]);
			}
			sb.append("#");
			sb.append(System.lineSeparator());
			String border = "##";
			for (int i = this.AA.length - 1; i >= 0; i--) {
				sb.append(border).append("#").append(this.AA[i]).append("#").append(this.BB[i]).append("#").append(this.CC[i]).append("#")
						.append(this.DD[i]).append("#").append(border).append(System.lineSeparator());
				border = "  ";
			}
			sb.append("  #########  ").append(System.lineSeparator());
			sb.append("  Price: " + this.price + "  ").append(System.lineSeparator());

			System.out.println(sb);
		}

		public boolean isDone() {
			return containsOnly(this.AA, A) && containsOnly(this.BB, B) && containsOnly(this.CC, C) && containsOnly(this.DD, D);
		}

		private boolean contains(final String[] toCheck, final String element) {
			for (final String current : toCheck) {
				if (StringUtils.equals(current, element)) {
					return true;
				}
			}
			return false;
		}

		private boolean containsOnly(final String[] toCheck, final String element) {
			for (final String current : toCheck) {
				if (!StringUtils.equals(current, element)) {
					return false;
				}
			}
			return true;
		}

		private boolean containsOnlyOrEmpty(final String[] toCheck, final String element) {
			for (final String current : toCheck) {
				if (!(StringUtils.equals(current, element) || StringUtils.equals(current, S))) {
					return false;
				}
			}
			return true;
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final GameState gameState = (GameState) o;

			// Probably incorrect - comparing Object[] arrays with Arrays.equals
			if (!Arrays.equals(this.AA, gameState.AA)) return false;
			// Probably incorrect - comparing Object[] arrays with Arrays.equals
			if (!Arrays.equals(this.BB, gameState.BB)) return false;
			// Probably incorrect - comparing Object[] arrays with Arrays.equals
			if (!Arrays.equals(this.CC, gameState.CC)) return false;
			// Probably incorrect - comparing Object[] arrays with Arrays.equals
			if (!Arrays.equals(this.DD, gameState.DD)) return false;
			// Probably incorrect - comparing Object[] arrays with Arrays.equals
			return Arrays.equals(this.HALL, gameState.HALL);
		}

		@Override public int hashCode() {
			int result = Arrays.hashCode(this.AA);
			result = 31 * result + Arrays.hashCode(this.BB);
			result = 31 * result + Arrays.hashCode(this.CC);
			result = 31 * result + Arrays.hashCode(this.DD);
			result = 31 * result + Arrays.hashCode(this.HALL);
			return result;
		}
	}

}
