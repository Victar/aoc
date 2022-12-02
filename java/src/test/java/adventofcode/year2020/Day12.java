package adventofcode.year2020;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day12 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day12/input.txt");
		System.out.println(solveSilver(data));
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day12/input.txt");
		System.out.println(solveGold(data));
	}

	public int solveGold(final ArrayList<String> data) {
		int x = 0;
		int y = 0;
		int waypointX = 10;
		int waypointY = 1;

		for (int i = 0; i < data.size(); i++) {
			final String instruction = data.get(i);
			final String command = instruction.substring(0, 1);
			final int value = Integer.parseInt(instruction.substring(1));
			if ("N".equals(command)) {
				waypointY = waypointY + value;
			}
			if ("S".equals(command)) {
				waypointY = waypointY - value;
			}
			if ("E".equals(command)) {
				waypointX = waypointX + value;
			}
			if ("W".equals(command)) {
				waypointX = waypointX - value;
			}
			if ("L".equals(command) || "R".equals(command)) {
				int direction = 0; //0-E, 1-S, 2-W, 3-N
				final int valuePour = (value % 360) / 90;
				if ("L".equals(command)) {
					direction = (direction - valuePour + 4) % 4;
				} else { //R case
					direction = (direction + valuePour + 4) % 4;
				}
				final int waypointXTemp = waypointX;
				final int waypointYTemp = waypointY;

				if (direction == 1) {
					waypointX = waypointYTemp;
					waypointY = -waypointXTemp;
				}
				if (direction == 2) {
					waypointX = -waypointXTemp;
					waypointY = -waypointYTemp;
				}
				if (direction == 3) {
					waypointX = -waypointYTemp;
					waypointY = waypointXTemp;
				}
			}
			if ("F".equals(command)) {
				x += waypointX * value;
				y += waypointY * value;
			}
		}
		return Math.abs(x) + Math.abs(y);
	}

	public int solveSilver(final ArrayList<String> data) {
		int x = 0;
		int y = 0;
		int direction = 0; //0-E, 1-S, 2-W, 3-N
		for (int i = 0; i < data.size(); i++) {
			final String instruction = data.get(i);
			final String command = instruction.substring(0, 1);
			final int value = Integer.parseInt(instruction.substring(1));
			//			System.out.println(command + "-" + value);
			if ("N".equals(command)) {
				y = y + value;
			}
			if ("S".equals(command)) {
				y = y - value;
			}
			if ("E".equals(command)) {
				x = x + value;
			}
			if ("W".equals(command)) {
				x = x - value;
			}
			if ("L".equals(command)) {
				System.out.println("L: " + value + "direction: " + direction);
				final int valuePour = (value % 360) / 90;
				direction = (direction - valuePour + 4) % 4;
			}
			if ("R".equals(command)) {
				System.out.println("R: " + value + "direction: " + direction);
				final int valuePour = (value % 360) / 90;
				direction = (direction + valuePour) % 4;
			}
			if ("F".equals(command)) {
				if (direction == 0) {
					x += value;
				}
				if (direction == 1) {
					y -= value;
				}
				if (direction == 2) {
					x -= value;
				}
				if (direction == 3) {
					y += value;
				}
			}
		}
		return Math.abs(x) + Math.abs(y);
	}
}

