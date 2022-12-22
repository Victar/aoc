package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day22 extends BaseTest {

	public static final int DAY = 22;

	public static int SIZE_X = Integer.MIN_VALUE;
	public static int SIZE_Y = Integer.MIN_VALUE;
	Map<String, Point> map;
	Point position = Point.of(0, 0, "x");

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		String instruction = data.get(data.size() - 1);
		for (int i = 0; i < data.size() - 1; i++) {
			SIZE_X = Math.max(SIZE_X, data.get(i).length());
		}
		SIZE_Y = data.size() - 2;
		System.out.println("SIZE_X:" + SIZE_X + " SIZE_Y:" + SIZE_Y);
		System.out.println(instruction);
		map = new HashMap<>();
		for (int y = 0; y < data.size() - 2; y++) {
			for (int x = 0; x < SIZE_X; x++) {
				String name = " ";
				if (data.get(y).length() > x) {
					name = data.get(y).charAt(x) + "";
				}
				Point current = new Point(x, y, name);
				map.put(current.getId(), current);
			}
		}
		Point direction = Point.FORWARD;
		String current = StringUtils.EMPTY;
		int i = 0;
		for (; i < instruction.length(); i++) {
			char cur = instruction.charAt(i);
			if (cur == 'R' || cur == 'L') {
				doInstruction(direction, Integer.parseInt(current));
				//				System.out.println(cur);
				direction = Point.getNextInstruction(direction, cur);
				//				System.out.println(Integer.parseInt(current));
				current = StringUtils.EMPTY;
			} else {
				current = current + cur;
			}
		}
		doInstruction(direction, Integer.parseInt(current));
		int bonusDirection = 0;
		if (Point.DOWN == direction) {
			bonusDirection = 1;
		}
		if (Point.BACK == direction) {
			bonusDirection = 2;
		}
		if (Point.UP == direction) {
			bonusDirection = 3;
		}
		System.out.println(position + "  " + direction);
		System.out.println(4 * (position.x + 1) + 1000 * (position.y + 1) + bonusDirection);
		drawState();
	}

	public void doInstruction(Point instruction, int steps) {
		//		System.out.println(steps + "  " + instruction.type  + " " + position);
		//		drawState();
		int cur = 0;
		Point lastField = position;
		for (; cur < steps; ) {
			Point next = position.nextPoint(instruction);
			Point nextOnMap = map.get(next.getId());
			if (nextOnMap.isWall()) {
				position = lastField;
				break;
			}
			if (nextOnMap.isSpace()) {
				position = next;
				continue;
			}
			if (nextOnMap.isField()) {
				cur++;
				position = next;
				lastField = position;

			}
		}
		//		drawState();
	}

	public void drawState() {
		System.out.println("      ");
		System.out.println(position);
		System.out.println("      ");

		for (int y = 0; y < SIZE_Y; y++) {
			StringBuilder sb = new StringBuilder();
			for (int x = 0; x < SIZE_X; x++) {
				String key = x + ":" + y;
				if (key.equals(position.getId())) {
					sb.append("x");
				} else {
					sb.append(map.get(key).type);
				}
				//				if (map.containsKey(key)) {
				//					sb.append(map.get(key).type);
				//				} else {
				//					System.out.println(key);
				//				}
			}
			System.out.println(sb);
		}
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

	@Data static class Point {

		public static Point FORWARD = Point.of(1, 0, "FORWARD");
		public static Point BACK = Point.of(-1, 0, "BACK");
		public static Point DOWN = Point.of(0, 1, "DOWN");
		public static Point UP = Point.of(0, -1, "UP");

		public static List<Point> INSTRUCTIONS = Arrays.asList(FORWARD, DOWN, BACK, UP);
		int x;
		int y;
		String type; // "#", ".", " ";
		public Point(int x, int y, String type) {
			this.x = x;
			this.y = y;
			this.type = type;
		}

		public static Point getNextInstruction(Point current, char direction) {
			int index = INSTRUCTIONS.indexOf(current);
			if ('R' == direction) {
				index = index + 1;
			}
			if ('L' == direction) {
				index = index - 1;
			}
			return INSTRUCTIONS.get((index + 4) % 4);
		}

		public static Point of(int x, int y) {
			return new Point(x, y, "");
		}

		public static Point of(int x, int y, String type) {
			return new Point(x, y, type);
		}

		public boolean isSpace() {
			return " ".equals(this.type);
		}

		public boolean isWall() {
			return "#".equals(this.type);
		}

		public boolean isField() {
			return ".".equals(this.type);
		}

		public Point nextPoint(Point direction) {
			int xNew = (this.x + direction.x + SIZE_X) % SIZE_X;
			int yNew = (this.y + direction.y + SIZE_Y) % SIZE_Y;
			return Point.of(xNew, yNew, this.type);
		}

		public String getId() {
			return x + ":" + y;
		}

		@Override public boolean equals(Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;
			Point point = (Point) o;
			return x == point.x && y == point.y;
		}

		@Override public int hashCode() {
			return Objects.hash(x, y);
		}

		@Override public String toString() {
			return "{" + "x=" + x + ", y=" + y + ", type='" + type + '\'' + '}';
		}
	}
}
